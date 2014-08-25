package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
)

func main() {
	orm, err := xorm.NewEngine("postgres", "postgres://postgres:postgres@localhost:5432/diel?sslmode=disable")
	orm.ShowDebug = true
	orm.ShowSQL = true

	if err != nil {
		fmt.Println(err)
		return
	}

	// Sync DB
	err = orm.Sync(new(Student))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer orm.Close()
	m := martini.Classic()
	m.Map(orm)
	m.Use(render.Renderer(render.Options{
		HTMLContentType: "application/json",
	}))
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "OPTIONS", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//Root
	m.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, "public/index.html")
	})

	//Student
	m.Get("/students", Students)
	m.Post("/students", binding.Json(StudentJSON{}), CreateStudent)
	m.Put("/students/:id", binding.Json(StudentJSON{}), UpdateStudent)
	m.Delete("/students/:id", binding.Json(StudentJSON{}), DeleteStudent)
	m.Run()
}
