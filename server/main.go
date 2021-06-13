package main

import (
	"fmt"   /*implements formatted I/O */
	"log"   /* implements a simple logging package. It defines a type, Logger, with methods for formatting output */
	"net"   /* provides a portable interface for network I/O */

	"grpc-server/chat"

	"google.golang.org/grpc"  /*install grpc go package to fetch neccessary dependencies */
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))   /* Sprintf() function in Go language formats according to a format specifier and returns the resulting string.*/
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()    /* NewServer creates a gRPC server which has no service registered and has not started to accept requests yet. */

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)  /* Fatalf is equivalent to Printf() followed by a call to os.Exit(1).*/
	}
}
