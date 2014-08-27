package main

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
)

type Student struct {
	Id           int64     `form:"id" json:"id"`
	Name         string    `form:"name" json:"name"`
	Registration string    `form:"registration" json:"registration"`
	CreatedAt    time.Time `form:"created_at" json:"created_at"`
	UpdatedAt    time.Time `form:"updated_at" json:"updated_at"`
}

type StudentJSON struct {
	Student Student `json:"student"`
}

func Students(db *gorm.DB, r render.Render) {
	students := make([]Student, 0)
	db.Find(&students)
	r.JSON(200, map[string]interface{}{"students": students})
}

func CreateStudent(db *gorm.DB, r render.Render, studentJson StudentJSON) {
	// student := studentJson.Student
	// _, err := orm.Insert(student)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// r.JSON(200, map[string]interface{}{"student": student})
}

// func UpdateStudent(orm *xorm.Engine, r render.Render, params martini.Params, studentJson StudentJSON) {
// 	student := studentJson.Student
// 	student.Id, _ = strconv.ParseInt(params["id"], 10, 64)
// 	_, err := orm.Id(student.Id).Update(&student)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	r.JSON(200, map[string]interface{}{"student": student})
// }

// func DeleteStudent(orm *xorm.Engine, r render.Render, params martini.Params) {
// 	_, err := orm.Id(params["id"]).Delete(&Student{})
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	r.JSON(204, map[string]interface{}{})
// }
