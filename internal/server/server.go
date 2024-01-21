package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"theobourgeois.com/internal/router"

	"github.com/joho/godotenv"
)

func Start() {
	router.SetupRoutes()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	fmt.Println("Server starting on port", port, "...")
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatalln(err)
		}
	}()

	select {}
}
