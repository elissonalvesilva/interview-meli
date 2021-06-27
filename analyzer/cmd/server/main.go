package main

import (
	"github.com/elissonalvesilva/interview-meli/analyzer/servers"
	"github.com/elissonalvesilva/interview-meli/pkg/proto/dna_pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	analyzerServer := servers.NewAnalyzerServer()

	grpcServer := grpc.NewServer()
	dna_pb.RegisterDNAServiceAnalyzerServer(grpcServer, analyzerServer)

	address := "0.0.0.0:50051"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
