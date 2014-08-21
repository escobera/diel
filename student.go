package main

import (
	"fmt"
	"time"

	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	"github.com/martini-contrib/render"
)

type Student struct {
	Id           int64     `form:"id" json:"id"`
	Name         string    `form:"name" json:"name"`
	Registration string    `form:"registration" json:"registration"`
	Created      time.Time `form:"created_at" json:"created_at" xorm:"CREATED"`
}

type StudentJSON struct {
	Student Student `json:"student"`
}

func Students(orm *xorm.Engine, r render.Render) {
	students := make([]Student, 0)
	orm.Find(&students)
	r.JSON(200, map[string]interface{}{"students": students})
}

func CreateStudent(orm *xorm.Engine, r render.Render, studentJson StudentJSON) {
	student := studentJson.Student
	_, err := orm.Insert(student)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.JSON(200, map[string]interface{}{"student": student})
}

func UpdateStudent(orm *xorm.Engine, r render.Render, params martini.Params, studentJson StudentJSON) {
	student := studentJson.Student
	fmt.Println(student)
	_, err := orm.Id(params["id"]).Update(student)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.JSON(200, map[string]interface{}{"student": student})
}

func DeleteStudent(orm *xorm.Engine, r render.Render, params martini.Params) {
	_, err := orm.Id(params["id"]).Delete(&Student{})
	if err != nil {
		fmt.Println(err)
		return
	}
	r.JSON(204, map[string]interface{}{})
}
