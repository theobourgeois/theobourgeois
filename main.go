package main

import (
	"theobourgeois.com/app/routes"
	"theobourgeois.com/internal/db"
	"theobourgeois.com/internal/server"
)

func main() {
	db.InitDB()
	routes.InitRoutes(db.GetDB())
	server.Start()
}
