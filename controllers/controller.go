package controllers

import (
	"github.com/Bobby-P-dev/assignment3-024/database"
	"github.com/Bobby-P-dev/assignment3-024/models"
	"github.com/gin-gonic/gin"
)

func UpdateWeather(c *gin.Context) {
	db := database.GetDB()
	micro := models.Microservice{}

	id := c.Param("id")

	data := db.First(&micro, id)

	if data.RowsAffected == 0 {
		createData := models.Microservice{
			ID:    1,
			Water: 0,
			Wind:  0}
		db.Create(&createData)
	} else {
		var newData models.Microservice
		c.BindJSON(&newData)
		micro.Water = newData.Water
		micro.Wind = newData.Wind
		db.Save(&micro)
	}
}
