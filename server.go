package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
)

func main() {
	orm, err := xorm.NewEngine("postgres", "postgres://rafa:@localhost:5432/diel?sslmode=disable")
	orm.ShowDebug = true
	orm.ShowSQL = true

	if err != nil {
		fmt.Println(err)
		return
	}

	// Sync DB
	err = orm.Sync(new(Estudante))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer orm.Close()
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		HTMLContentType: "application/json",
	}))
	m.Map(orm)

	//Root
	m.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, "public/index.html")
	})

	//Estudante
	m.Get("/estudantes", Estudantes)
	m.Post("/estudantes", binding.Json(EstudanteJSON{}), CreateEstudante)
	m.Patch("/estudantes/:id", binding.Json(EstudanteJSON{}), UpdateEstudante)
	m.Delete("/estudantes/:id", binding.Json(EstudanteJSON{}), DeleteEstudante)

	m.Run()
}
