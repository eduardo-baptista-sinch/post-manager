package main

import "main/cmd/post_manager/http"

func main() {
	server := http.NewServer()
	server.Listen()
}
