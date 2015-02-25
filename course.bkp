// package main
//
// import (
// 	"fmt"
// 	"strconv"
// 	"time"
//
// 	"github.com/go-martini/martini"
// 	"github.com/jinzhu/gorm"
// 	"github.com/martini-contrib/render"
// )
//
// type Course struct {
// 	Id         int64     `form:"id" json:"id"`
// 	Name       string    `form:"name" json:"name"`
// 	Sequential int8      `form:"sequential" json:"sequential"`
// 	Year       int8      `form:"year" json:"year"`
// 	CreatedAt  time.Time `form:"created_at" json:"created_at"`
// 	UpdatedAt  time.Time `form:"updated_at" json:"updated_at"`
// 	Students   []Student
// }
//
// // type CourseJSON struct {
// // 	Student Student `json:"course"`
// // }
//
// // func Courses(orm *xorm.Engine, r render.Render) {
// // 	students := make([]Courses, 0)
// // 	orm.Find(&students)
// // 	r.JSON(200, map[string]interface{}{"students": students})
// // }
//
// // func CreateStudent(orm *xorm.Engine, r render.Render, studentJson StudentJSON) {
// // 	student := studentJson.Student
// // 	_, err := orm.Insert(student)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}
// // 	r.JSON(200, map[string]interface{}{"student": student})
// // }
//
// // func UpdateStudent(orm *xorm.Engine, r render.Render, params martini.Params, studentJson StudentJSON) {
// // 	student := studentJson.Student
// // 	student.Id, _ = strconv.ParseInt(params["id"], 10, 64)
// // 	_, err := orm.Id(student.Id).Update(&student)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}
// // 	r.JSON(200, map[string]interface{}{"student": student})
// // }
//
// // func DeleteStudent(orm *xorm.Engine, r render.Render, params martini.Params) {
// // 	_, err := orm.Id(params["id"]).Delete(&Student{})
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}
// // 	r.JSON(204, map[string]interface{}{})
// // }
