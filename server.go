package main

import (
	_ "encoding/json"

	"github.com/gin-gonic/gin"
)

func main() {
	// db, err := gorm.Open("postgres", "user=rafa dbname=diel sslmode=disable")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer db.Close()
	// db.LogMode(true)
	// // Migrate DB
	// db.AutoMigrate(Student{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	router := gin.Default()

	router.Get("/", func(c *gin.Context) {
		c.HTML(200, "public/index.html")
	})

	router.Run(":3001")

	// m.Map(&db)
	// m.Use(render.Renderer(render.Options{
	// 	HTMLContentType: "application/json",
	// }))
	// m.Use(cors.Allow(&cors.Options{
	// 	AllowOrigins:     []string{"http://localhost:4200"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "POST", "OPTIONS", "GET", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	//Root
	// router.Get("/", func(c *gin.Context) {
	// 	w.Header().Set("Content-Type", "text/html")
	// 	http.ServeFile(w, r, "public/index.html")
	// })

	//Student
	// m.Get("/students", Students)
	// m.Post("/students", binding.Json(StudentJSON{}), CreateStudent)
	// m.Put("/students/:id", binding.Json(StudentJSON{}), UpdateStudent)
	// m.Delete("/students/:id", binding.Json(StudentJSON{}), DeleteStudent)

	//Course
	// m.Get("/courses", Courses)
	// m.Post("/courses", binding.Json(CourseJSON{}), CreateCourse)
	// m.Put("/courses/:id", binding.Json(CourseJSON{}), UpdateCourse)
	// m.Delete("/courses/:id", binding.Json(CourseJSON{}), DeleteCourse)

}
