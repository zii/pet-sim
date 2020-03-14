package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"

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
	Render(w, "index.html", pongo2.Context{})
}

func r_newpet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	no := base.ToInt(vars["id"])
	fmt.Println("no:", no)
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
	fmt.Println("newchar:", char.Name, char.Lv, char.WorkMaxHp, char.WorkFixStr, char.WorkFixTough, char.WorkFixDex)
	Render(w, "newpet.html", pongo2.Context{"pet": char})
}

func r_pet(w http.ResponseWriter, r *http.Request) {
}
