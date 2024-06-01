package main

import (
	"GO/database"
	"GO/router"
)

func main() {
	database.ConnectDB()
	r := router.SetupRouter()
	r.Run(":8080")
}