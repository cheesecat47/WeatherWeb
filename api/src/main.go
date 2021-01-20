package main

import (
	"app"
	"os"
)

func main() {
	app.Start(os.Getenv("port"))
}
