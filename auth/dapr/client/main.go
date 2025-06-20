package main

import (
	"context"
	"log"

	pb "auth/proto/greeterpb"

	dapr "github.com/dapr/go-sdk/client"
	"google.golang.org/protobuf/proto"
)

func main() {
	ctx := context.Background()

	client, err := dapr.NewClient()
	if err != nil {
		log.Fatalf("failed to create Dapr client: %v", err)
	}
	defer client.Close()

	req := &pb.HelloRequest{Name: "Mihir"}
	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatalf("failed to marshal proto: %v", err)
	}

	respBytes, err := client.InvokeMethodWithContent(
		ctx,
		"auth",
		"proto.Greeter/SayHello",
		"application/x-protobuf",
		&dapr.DataContent{
			ContentType: "application/x-protobuf",
			Data:        data,
		},
	)
	if err != nil {
		log.Fatalf("invoke failed: %v", err)
	}

	var reply pb.HelloReply
	err = proto.Unmarshal(respBytes, &reply)
	if err != nil {
		log.Fatalf("unmarshal failed: %v", err)
	}

	log.Printf("Reply from auth: %s", reply.GetMessage())
}
