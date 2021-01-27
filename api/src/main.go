package main

import (
	"github.com/cheesecat47/webpractice/api/constant"
	"github.com/cheesecat47/webpractice/api/server"
)

func main() {
	server.Start(constant.ListeningPort)
}
