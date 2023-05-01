package main

import (
	"awesomeProject/src/httpserver"
	"awesomeProject/src/ws"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// Load the environment file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load the env file.", err)
	}
}

func main() {
	server := flag.String("server", "", "http,websocket")
	flag.Parse()
	switch *server {
	case "http":
		fmt.Println("http server is starting on :8080")
		httpserver.StartHTTPServer()
	case "websocket":
		fmt.Println("websocket server is starting on :8081")
		ws.StartWebsocketServer()
	default:
		fmt.Println("invalid server. Available server: http or websocket")
	}
}
