package main

import (
	"github.com/errolitolopez/geo-api/db"
	"github.com/errolitolopez/geo-api/routes"
)

func main() {
	// Database setup
	db := db.SetupDB()
	// Routes setup
	routes := routes.SetupRoutes(db)
	routes.Run("localhost:9999")
}
