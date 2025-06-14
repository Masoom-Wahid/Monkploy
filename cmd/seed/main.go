package main

import (
	"fmt"
	"os"

	"platform/database/seeder"
)

var seeders = map[string]func(){
	"apps": func() {
		seeder.NewSeederSupplier().AppSeeder().SeedApps()
	},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/seed/main.go <seederName>|all")
		os.Exit(1)
	}

	seederArgs := os.Args[1:]

	if len(seederArgs) == 1 && seederArgs[0] == "all" {
		for name, fn := range seeders {
			fmt.Printf("Running seeder: %s\n", name)
			fn()
		}
		return
	}

	for _, seederName := range seederArgs {
		fn, ok := seeders[seederName]
		if !ok {
			fmt.Printf("Unknown seeder: %s\n", seederName)
			continue
		}
		fmt.Printf("Running seeder: %s\n", seederName)
		fn()
	}
}
