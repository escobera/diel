package service

import (
	"fmt"

	"github.com/escobera/diel/resource"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Config struct {
	SvcHost    string
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
}

type Diel struct {
}

func (s *Diel) getDb(cfg Config) (gorm.DB, error) {
	connectionString := fmt.Sprintf("user=%s dbname=%s sslmode=disabled", cfg.DbUser, cfg.DbPassword)

	return gorm.Open("postgres", connectionString)
}

func (s *Diel) Migrate(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.LogMode(true)
	defer db.Close()

	db.AutoMigrate(&resource.Student{})
	return nil
}
func (s *Diel) Run(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	// r.GET("/todo", todoResource.GetAllTodos)
	// r.GET("/todo/:id", todoResource.GetTodo)
	// r.POST("/todo", todoResource.CreateTodo)
	// r.PUT("/todo/:id", todoResource.UpdateTodo)
	// r.PATCH("/todo/:id", todoResource.PatchTodo)
	// r.DELETE("/todo/:id", todoResource.DeleteTodo)

	r.Run(cfg.SvcHost)

	return nil
}
