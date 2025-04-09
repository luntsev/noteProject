package main

import "noteProject/server"

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
}
