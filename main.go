package main

import (
	"fmt"
	"go-auth/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/send", controllers.Send).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://promise-jeremiah-portfolio.herokuapp.com/"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	fmt.Println("Server Listening ...")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
