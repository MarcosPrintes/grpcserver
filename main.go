package main

import (
	"fmt"
	"log"
	"net"

	proto_buff "github.com.br/MarcosPrintes/grpc/customer"
	"github.com.br/MarcosPrintes/grpc/server"

	"google.golang.org/grpc"
)

var serv server.Server

const (
	port = ":8080"
)

func main() {
	fmt.Println("start server")
	serv = server.Server{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer()
	proto_buff.RegisterFakeCustomerServer(s, &serv) //Register service to communicate with client
	s.Serve(lis)
}
