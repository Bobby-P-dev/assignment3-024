package main

import (
	"github.com/Bobby-P-dev/assignment3-024/database"
	"github.com/Bobby-P-dev/assignment3-024/routes"
	"github.com/Bobby-P-dev/assignment3-024/services"
)

func main() {
	database.ConnectToDB()

	go services.HitAPI()

	router := routes.SetupRouter()

	router.Run(":8080")
}
