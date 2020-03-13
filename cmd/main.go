package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/zii/pet-sim/base"
)

var args struct {
	Listen string
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
	flag.StringVar(&args.Listen, "listen", ":80", "server http listening address")
}

func main() {
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", r_index)
	router.HandleFunc("/pet/{id}", r_pet)
	router.PathPrefix("/f/").Handler(http.StripPrefix("/f/", http.FileServer(http.Dir("static"))))

	// Interrupt handler.
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Listen http.
	go func() {
		log.Println("Stoneage pet simulator, server listening http", args.Listen)
		s := &http.Server{
			Addr:        args.Listen,
			Handler:     router,
			IdleTimeout: 60 * time.Second,
		}
		err := s.ListenAndServe()
		base.Raise(err)
	}()

	// Run!
	log.Println("exit", <-errc)
}
