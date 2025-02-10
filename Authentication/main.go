package main

import (
	"context"
	"fmt"
	"log"
	"main/database"
	"net/http"

	"github.com/conglt10/web-golang/routes"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	db := database.ConnectDB()
	defer db.Close(context.Background())

	router := httprouter.New()
	router.POST("/auth/login",routes.Login)
	router.POST("/auth/register", routes.Register)

	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}