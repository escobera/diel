package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
	"net/http"
	"time"
)

type Estudante struct {
	Id        int64     `form:"id" json:"id"`
	Nome      string    `form:"nome" json:"nome"`
	Matricula string    `form:"matricula" json:"matricula"`
	Created   time.Time `form:"created_at" json:"created_at" xorm:"created"`
}

func Estudantes(orm *xorm.Engine, r render.Render) {
	var estudantes []Estudante
	orm.Find(&estudantes)
	r.JSON(200, map[string]interface{}{"estudantes": estudantes})
}

func CreateEstudante(orm *xorm.Engine, r render.Render, estudante Estudante, req *http.Request) {
	//estudante := new(Estudante)
	//	estudante.Nome = params["nome"]
	//_, err := orm.Insert(estudante)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(estudante)
	//r.JSON(200, map[string]interface{}{"estudante": estudante})
	fmt.Println(estudante)
	return
}

func main() {
	orm, err := xorm.NewEngine("postgres", "postgres://postgres:postgres@localhost:5432/diel?sslmode=disable")
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

	m.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, "public/index.html")
	})

	m.Get("/estudantes", Estudantes)
	m.Post("/estudantes", binding.Json(Estudante{}), CreateEstudante)
	m.Run()
}
