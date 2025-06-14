package main

import (
	"fmt"
	"platform/database"
	"platform/database/migrations"
)


func main() {
	db := database.Connect()

	fmt.Println("Running Migrations....")
	migrations.Migrate(db)
}