package server

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"

	proto_buff "github.com.br/MarcosPrintes/grpc/customer"
)

// Server is used to implement customer.CustomerServer.
type Server struct {
	savedFakersCustomers []*proto_buff.FakeCustomerRequest
}

//CreateFakeCustomer bla bla
func (server *Server) CreateFakeCustomer(ctx context.Context, fakeRequest *proto_buff.FakeCustomerRequest) (*proto_buff.FakeResponse, error) {
	fmt.Println("server create customer")
	fmt.Println("server receive request => ", fakeRequest)

	server.savedFakersCustomers = append(server.savedFakersCustomers, fakeRequest)
	return &proto_buff.FakeResponse{Id: fakeRequest.Id, Success: true}, nil
}

//GetFakersCustomer bla bla
func (server *Server) GetFakersCustomer(filter *proto_buff.FakeCustomerFilter, stream proto_buff.FakeCustomer_GetFakersCustomerServer) error {
	for _, fakeCustomer := range server.savedFakersCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(fakeCustomer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(fakeCustomer); err != nil {
			return err
		}
	}
	return nil
}

//CustomMessage bla bla
func (server *Server) SendCustomMessage(ctx context.Context, customMessage *proto_buff.CustomMessageRequest) (*proto_buff.MessageResponse, error) {
	fmt.Println(customMessage)
	return &proto_buff.MessageResponse{Id: 1, Success: true}, nil
}
