package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type StudentApi struct {
	db gorm.DB
}

func (api *StudentApi) CreateStudent(c *gin.Context) {
	var student resource.Student

	c.Bind(&student)

	api.db.Save(&student)
	c.JSON(201, student)
}
