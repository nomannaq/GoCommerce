package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/nomannaq/webshop-api-golang/cmd/controllers"
	config "github.com/nomannaq/webshop-api-golang/cmd/database"
	"github.com/nomannaq/webshop-api-golang/cmd/models"
)

func init() {
	//Load the .env file
	if err := godotenv.Load("./db.env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func main() {

	config.InitDB()
	config.DB.AutoMigrate(&models.Product{}, &models.User{})

	//Routes
	http.HandleFunc("/register", controllers.RegisterUser)
	http.HandleFunc("/login", controllers.LoginUser)
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
