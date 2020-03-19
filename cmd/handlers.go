package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"

	"github.com/zii/pet-sim/biz"

	"github.com/zii/pet-sim/base"

	"github.com/flosch/pongo2"
)

var TplSet *pongo2.TemplateSet

func init() {
	loader := pongo2.MustNewLocalFileSystemLoader("tpl")
	TplSet = pongo2.NewSet("set1", loader)
}

func Render(w io.Writer, path string, args pongo2.Context) {
	t := pongo2.Must(TplSet.FromFile(path))
	err := t.ExecuteWriter(args, w)
	base.Raise(err)
}

func r_index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tpl/index.html")
	//Render(w, "index.html", pongo2.Context{"pet": char})
}

func api_getpet(w http.ResponseWriter, r *http.Request) {
	var args = struct {
		No int `json:"no"`
	}{}
	dec := json.NewDecoder(r.Body)
	dec.Decode(&args)

	var no int
	if args.No != 0 {
		no = args.No
	} else {
		i := rand.Intn(len(biz.EnemyNoList))
		no = biz.EnemyNoList[i]
	}
	eb := biz.GetEnemyBase(no)
	if eb == nil {
		http.NotFound(w, r)
		return
	}
	char := biz.CreateEnemy(no, 1)
	if char == nil {
		http.NotFound(w, r)
		return
	}
	w.Header().Add("Conent-Type", "application/json")
	b, _ := json.Marshal(char)
	w.Write(b)
}

func api_levelup(w http.ResponseWriter, r *http.Request) {
	var args = struct {
		Id int `json:"id"`
		Up int `json:"up"`
	}{}
	dec := json.NewDecoder(r.Body)
	dec.Decode(&args)
	pet := biz.GetChar(args.Id)
	if pet == nil {
		http.NotFound(w, r)
		return
	}
	for i := 0; i < args.Up; i++ {
		biz.PetLevelUp(pet)
	}
	w.Header().Add("Conent-Type", "application/json")
	b, _ := json.Marshal(pet)
	w.Write(b)
}
