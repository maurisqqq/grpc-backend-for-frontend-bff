package main

import (
	"context"
	"log"

	pb "grpc-unary/hello/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr                  string = "localhost:50051"
	userGrpcServiceClient pb.HelloServiceClient
)

func prepareGrpcClient(ctx context.Context) error {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	if userGrpcServiceClient != nil {
		conn.Close()
		return nil
	}
	userGrpcServiceClient = pb.NewHelloServiceClient(conn)

	return nil
}

func doHello(ctx context.Context, req string) (*pb.HelloResponse, error) {
	log.Println("doHello was invoked")
	if err := prepareGrpcClient(ctx); err != nil {
		return nil, err
	}

	r, err := userGrpcServiceClient.Hello(context.Background(), &pb.HelloRequest{FirstName: req})

	if err != nil {
		log.Fatalf("Could not say hello: %v\n", err)
	}

	return r, nil
}
