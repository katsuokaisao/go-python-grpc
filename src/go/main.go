package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/katsuokaisao/proto-gen-sample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 1234
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterPinPonServiceServer(s, NewPinPonSever())

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

func NewPinPonSever() *PinPonSever {
	return &PinPonSever{}
}

type PinPonSever struct {
	pb.UnimplementedPinPonServiceServer
}

func (s *PinPonSever) Send(ctx context.Context, req *pb.PinPonRequest) (*pb.PinPonResponse, error) {
	fmt.Printf("receive %s\n", req.GetWord())
	return &pb.PinPonResponse{Word: "pong"}, nil
}
