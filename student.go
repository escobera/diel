package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"strconv"
	"time"
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
	student := studentJson.Student
	err := db.Create(&student)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.JSON(200, map[string]interface{}{"student": student})
}

func UpdateStudent(db *gorm.DB, r render.Render, params martini.Params, studentJson StudentJSON) {
	student := studentJson.Student
	student.Id, _ = strconv.ParseInt(params["id"], 10, 64)
	err := db.Save(&student)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.JSON(200, map[string]interface{}{"student": student})
}

func DeleteStudent(db *gorm.DB, r render.Render, params martini.Params) {
	student := &Student{}
	student.Id, _ = strconv.ParseInt(params["id"], 10, 64)
	err := db.Delete(&student)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.JSON(204, map[string]interface{}{})
}
