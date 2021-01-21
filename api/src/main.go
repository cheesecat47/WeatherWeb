package main

import (
	"github.com/cheesecat47/webpractice/constant"
	"github.com/cheesecat47/webpractice/server"
)

func main() {
	server.Start(constant.ListeningPort)
}
