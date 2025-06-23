package main

import (
	"log"
	"techtestify/internal/db"
	"techtestify/internal/router"
)

func main() {
	db.Init()
	r := router.SetupRouter()

	log.Println("🚀 Starting server on :8080")
	r.Run(":8080")
}
