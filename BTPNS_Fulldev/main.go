package main

import (
	"log"

	"BTPNS_Fulldev/database"
	"BTPNS_Fulldev/router"
)

func main() {
    log.Println("Initializing database...")
    database.InitDB()
    log.Println("Database initialized.")

    r := router.SetupRouter()
    log.Println("Starting server on localhost:8080")
    if err := r.Run("localhost:8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
