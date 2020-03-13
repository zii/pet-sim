package main

import (
	"io"
	"net/http"

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

func r_pet(w http.ResponseWriter, r *http.Request) {
}
