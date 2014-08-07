package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type Estudante struct {
	Id        int64
	Nome      string
	Matricula string
	Created   time.Time `xorm:"created"`
}

func Estudantes(orm *xorm.Engine, params martini.Params) {
	var estudantes []Estudante
	orm.Find(&estudantes)
	resultado, _ := json.Marshal(estudantes)
	os.Stdout.Write(resultado)
}

func main() {
	orm, err := xorm.NewEngine("postgres", "postgres://rafa:@localhost:5432/diel?sslmode=disable")
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
	m.Map(orm)
	m.Get("/estudantes", Estudantes)
	m.Run()
}
