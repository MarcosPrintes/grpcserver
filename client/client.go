package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	proto_buff "github.com.br/MarcosPrintes/grpc/customer"
)

const (
	address = "localhost:8080"
)

func createFakeCustomer(client proto_buff.FakeCustomerClient, customer *proto_buff.FakeCustomerRequest) {
	fmt.Println("client create customer")
	resp, err := client.CreateFakeCustomer(context.Background(), customer)
	if err != nil {
		log.Fatal("Could not create a fake customer ", err.Error())
	}
	if resp.Success {
		log.Printf("Fake Customer Created %v", resp.Id)
	}
}

func getFakers(client proto_buff.FakeCustomerClient, filter *proto_buff.FakeCustomerFilter) {
	stream, err := client.GetFakersCustomer(context.Background(), filter)
	if err != nil {
		log.Fatal("Error on get fakers", err.Error())
	}

	for {
		fakeCustomers, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetFakeCustomers(_) = _, %v", client, err)
		}
		log.Printf("Faker %v", fakeCustomers)
	}
}

func sendCustomMessage(client proto_buff.FakeCustomerClient, customMessage *proto_buff.CustomMessageRequest) {
	fmt.Println("client  custom message => ", customMessage)
	response, err := client.SendCustomMessage(context.Background(), customMessage)
	if err != nil {
		log.Fatal("send message error =>", err.Error())
	}
	fmt.Println(response)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // CONNECTION TO gRPC SERVER.
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() // CLOSE CONNECTION WHEN END main()

	fakeClient := proto_buff.NewFakeCustomerClient(conn)

	fake := &proto_buff.FakeCustomerRequest{
		Id:    1,
		Name:  "Um Nome Aew",
		Phone: "(92) 9 9999-9999",
	}

	createFakeCustomer(fakeClient, fake)

	fakeFilter := &proto_buff.FakeCustomerFilter{Keyword: ""}
	getFakers(fakeClient, fakeFilter)
	msg := &proto_buff.CustomMessageRequest{
		TypeMsg: "message from client 2",
		Name:    "Teste Client 2",
	}
	sendCustomMessage(fakeClient, msg)
}
