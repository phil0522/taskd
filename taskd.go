//go:generate protoc --go_out=plugins=grpc:. proto/snippet_service.proto

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/phil0522/taskd/pkg/server"
	pb "github.com/phil0522/taskd/proto"
	"google.golang.org/grpc"
	"gopkg.in/sevlyar/go-daemon.v0"
)

const (
	serverAddr = "127.0.0.1:6398"
)

var debugLogging = false

func serve() {
	lis, err := net.Listen("tcp4", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	snipperServer := &server.SnippetServer{
		GrpcServer: grpcServer,
	}
	snipperServer.Initialize()
	pb.RegisterSnippetServiceServer(grpcServer, snipperServer)

	grpcServer.Serve(lis)
	log.Printf("server exit")
}

func formatSnippet(snippet *pb.ShellSnippet) {
	name := strings.ReplaceAll(snippet.GetSnippetName(), ":", "")
	desc := strings.ReplaceAll(snippet.GetSnippetDescription(), ":", "")
	cmd := strings.ReplaceAll(snippet.GetSnippetCommand(), "\n", "")
	fmt.Printf("%s:%s:%s\n", name, desc, cmd)
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
	for _, snippet := range resp.ShellSnippet {
		formatSnippet(snippet)
	}
}

func main() {
	userRoot := os.Getenv("HOME")
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
		time.Sleep(time.Second * 2)
		request()
	} else {
		serve()
		defer context.Release()
	}
}
