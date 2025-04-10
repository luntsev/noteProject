package main

import "noteproject/server"

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
}
