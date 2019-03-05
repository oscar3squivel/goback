package main

import (
	"log"
	"net/http"
	. "travel-n-expenses/api"
	DB "travel-n-expenses/config"

	"github.com/rs/cors"
)

func main() {
	DB.ConnectDataBase()
	//Loading all routes from api/index.gp
	router := AddAllRoutes()
	CORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})
	//Enable CORS, Allows domains to get access to Golang API
	handler := CORS.Handler(router)
	log.Fatal(http.ListenAndServe(":3001", handler))
}
