package main

import (
	"os"

	"./controllers"
	"./database"
	"./middlewares"
	"./routes"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv
	if port == "" {
		port = "8080"
	}
	app :=controllers.NewApplication(database.ProductData(database.Client, "products"), database.UserData(database.Client, "users"))	

}