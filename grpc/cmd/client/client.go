package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jhosefmarks/grpc-labs/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	fmt.Println("\n---- Call AddUser ----")
	AddUser(client)
	fmt.Println("\n---- Call AddUserVerbose ----")
	AddUserVerbose(client)
	fmt.Println("\n---- Call AddUsers ----")
	AddUsers(client)
	fmt.Println("\n---- Call AddUserStreamBoth ----")
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		// Id:    faker.UUIDDigit(),
		Name:  faker.Name(),
		Email: faker.Email(),
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    faker.UUIDDigit(),
		Name:  faker.Name(),
		Email: faker.Email(),
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{}

	for i := 1; i <= 5; i++ {
		reqs = append(reqs, &pb.User{
			Id:    faker.UUIDDigit(),
			Name:  faker.Name(),
			Email: faker.Email(),
		})
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{}

	for i := 1; i <= 5; i++ {
		reqs = append(reqs, &pb.User{
			Id:    faker.UUIDDigit(),
			Name:  faker.Name(),
			Email: faker.Email(),
		})
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}
			fmt.Printf("Receiving user %v with status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
