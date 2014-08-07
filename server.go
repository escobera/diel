package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
)

type Estudante struct {
	Id        int64     `json:"id"`
	Nome      string    `json:"nome"`
	Matricula string    `json:"matricula"`
	Created   time.Time `json:"created_at" xorm:"created"`
}

func Estudantes(orm *xorm.Engine, r render.Render, params martini.Params) {
	var estudantes []Estudante
	orm.Find(&estudantes)
	r.JSON(200, map[string]interface{}{"estudantes": estudantes})
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
	m.Run()
}
