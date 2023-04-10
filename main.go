package main

import (
	"project-1-chapter-2/database"
	"project-1-chapter-2/routers"
)

func main() {

	database.StartDB()

	routers.StartServer().Run(":8080")

}
