package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zii/pet-sim/biz"

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
	rand.Seed(time.Now().Unix())

	// init data
	biz.InitSkill("data/petskill2.txt")
	biz.InitMagic("data/magic.txt")
	biz.InitEnemyBase("data/enemybase2.txt")
	biz.InitChar()

	// http server
	router := mux.NewRouter()
	router.HandleFunc("/", r_index)
	router.PathPrefix("/f/").Handler(http.StripPrefix("/f/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/api/getpet", api_getpet)
	router.HandleFunc("/api/levelup", api_levelup)

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
