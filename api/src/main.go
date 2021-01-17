package main

import {
	app
}

func main() {
	app.Start(os.Getenv("port"));
}