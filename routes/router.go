package routes

import (
	"github.com/Bobby-P-dev/assignment3-024/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.PUT("/update/:id", controllers.UpdateWeather)

	return r

}
