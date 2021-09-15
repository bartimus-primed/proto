package main

import (
	"context"
	"log"
	"net"

	pb "github.com/bartimus-primed/proto/reverse/reverse_pb"
	"google.golang.org/grpc"
)

const (
	port = ":50551"
)

// server is used to implement Ghost Server.
type server struct {
	pb.UnimplementedReverseInteractServer
}

func (s *server) GetCommand(ctx context.Context, in *pb.Response) (*pb.Command, error) {
	if in.GetReady() {

	}
	return &pb.Command{}, nil
}

func (s *server) HandsOn(stream pb.ReverseInteract_HandsOnServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReverseInteractServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
