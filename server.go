package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=diel sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	db.LogMode(true)
	// Migrate DB
	db.AutoMigrate(Student{})
	if err != nil {
		fmt.Println(err)
		return
	}
	m := martini.Classic()
	m.Map(&db)
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

	//Course
	// m.Get("/courses", Courses)
	// m.Post("/courses", binding.Json(CourseJSON{}), CreateCourse)
	// m.Put("/courses/:id", binding.Json(CourseJSON{}), UpdateCourse)
	// m.Delete("/courses/:id", binding.Json(CourseJSON{}), DeleteCourse)
	m.Run()
}
