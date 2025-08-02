package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/thienchuong/golang-file-upload/internal/stream"
	"github.com/thienchuong/golang-file-upload/proto"
)

func main() {
	grpcServer := grpc.NewServer()

	streamingService := &stream.Service{}
	proto.RegisterFileUploadServiceServer(grpcServer, streamingService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("starting grpc server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
