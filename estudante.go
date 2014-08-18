package main

import (
	"fmt"
	"time"

	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	"github.com/martini-contrib/render"
)

type Estudante struct {
	Id        int64     `form:"id" json:"id"`
	Nome      string    `form:"nome" json:"nome"`
	Matricula string    `form:"matricula" json:"matricula"`
	Created   time.Time `form:"created_at" json:"created_at" xorm:"CREATED"`
}

type EstudanteJSON struct {
	Estudante Estudante `json:"estudante"`
}

func Estudantes(orm *xorm.Engine, r render.Render) {
	var estudantes []Estudante
	orm.Find(&estudantes)
	r.JSON(200, map[string]interface{}{"estudantes": estudantes})
}

func CreateEstudante(orm *xorm.Engine, r render.Render, estudanteJson EstudanteJSON) {
	estudante := estudanteJson.Estudante
	_, err := orm.Insert(estudante)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.JSON(200, map[string]interface{}{"estudante": estudante})
}

func UpdateEstudante(orm *xorm.Engine, r render.Render, params martini.Params, estudanteJson EstudanteJSON) {
	estudante := estudanteJson.Estudante
	fmt.Println(estudante)
	_, err := orm.Id(params["id"]).Update(estudante)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.JSON(200, map[string]interface{}{"estudante": estudante})
}

func DeleteEstudante(orm *xorm.Engine, r render.Render, params martini.Params) {
	_, err := orm.Id(params["id"]).Delete(&Estudante{})
	if err != nil {
		fmt.Println(err)
		return
	}
	r.JSON(204, map[string]interface{}{})
}
