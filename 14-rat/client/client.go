package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cleonty/rat/grpcapi"

	"google.golang.org/grpc"
)

var (
	opts   []grpc.DialOption
	conn   *grpc.ClientConn
	err    error
	client grpcapi.AdminClient
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: client cmd\n")
		os.Exit(0)
	}
	opts = append(opts, grpc.WithInsecure())
	if conn, err = grpc.Dial(fmt.Sprintf("localhost:%d", 9090), opts...); err != nil {

		log.Fatal(err)
	}

	defer conn.Close()
	log.Printf("client connected to admin server on port 9090\n")
	client = grpcapi.NewAdminClient(conn)
	var cmd = new(grpcapi.Command)
	cmd.In = os.Args[1]
	ctx := context.Background()
	fmt.Printf("client about to run command %s on admin server\n", cmd.In)
	cmd, err = client.RunCommand(ctx, cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("client got response from admin server\n")
	fmt.Println(cmd.Out)
}
