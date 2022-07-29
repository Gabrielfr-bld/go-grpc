package main 

import (
	"context"
	"log"

	"github.com/Gabrielfr-bld/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to dial: ", err)
	}

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Hello GRPC",
	}

	res, err := client.RequestMessage(context.Background(), req)
	if err != nil {
		log.Fatal("failed to request: ", err)
	}

	log.Print("response: ", res.GetStatus())
}