package main

import (
	"os"
	"platform/app"
	"time"
)

func main() {
	os.Setenv("TZ", "Asia/Kabul")
	loc, err := time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		panic(err)
	}

	// set the global timezone
	time.Local = loc

	app := app.NewApp()
	app.Run()
}
