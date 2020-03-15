//go:generate protoc --go_out=plugins=grpc:. proto/snippet_service.proto

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/phil0522/taskd/pkg/server"
	pb "github.com/phil0522/taskd/proto"
	"google.golang.org/grpc"
	"gopkg.in/sevlyar/go-daemon.v0"
)

const (
	serverAddr = "127.0.0.1:6398"
)

func serve() {
	lis, err := net.Listen("tcp4", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	snipperServer := &server.SnippetServer{
		GrpcServer: grpcServer,
	}
	pb.RegisterSnippetServiceServer(grpcServer, snipperServer)

	grpcServer.Serve(lis)
	log.Printf("server exit")
}

func request() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create connection, %s", err.Error())
	}

	client := pb.NewSnippetServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SearchShellSnippet(ctx, &pb.ShellSnippetRequest{})
	if err != nil {
		log.Fatalf("failed to execute rpc %s", err.Error())
	}
	log.Printf("get resposne %v", resp)
}

func main() {
	userRoot := os.Getenv("HOME")
	fmt.Printf("%s\n", userRoot)
	context := daemon.Context{
		PidFileName: "taskd.lock",
		PidFilePerm: 0644,
		LogFileName: "taskd.log",
		LogFilePerm: 0666,
		WorkDir:     userRoot,
	}

	child, _ := context.Search()
	if child != nil {
		log.Printf("Server has been already serving")
		request()
	}

	child, err := context.Reborn()
	if err != nil {
		log.Fatal(err.Error())
	}

	if child != nil {
		log.Printf("I am the parent, my pid %d and my child pid %d", os.Getpid(), child.Pid)
		time.Sleep(time.Second * 5)
		request()
	} else {
		serve()
		defer context.Release()
	}
}
