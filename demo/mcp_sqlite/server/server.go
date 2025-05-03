package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	mcpsqlite "github.com/peacess/go/demo/mcp_sqlite"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())
	serive := mcpsqlite.ServiceSqlite{}

	serive.Register(server)

	// Start the server
	log.Println("MCP Server is now running and waiting for requests...")
	err := server.Serve()
	if err != nil {
		println(err)
	} else {
		println("server started")
		<-interrupt
	}
	serive.Uninit()
	println("server exit")
}
