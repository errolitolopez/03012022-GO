package main

import (
	"github.com/errolitolopez/dynamo-db-geo-api/db"
	"github.com/errolitolopez/dynamo-db-geo-api/routes"
)

func main() {
	db := db.SetupDB()
	routes := routes.SetupRoutes(db)
	routes.Run("localhost:9999")
}
