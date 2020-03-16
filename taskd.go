//go:generate protoc --go_out=plugins=grpc:. proto/snippet_service.proto

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/phil0522/taskd/pkg/server"
	"github.com/phil0522/taskd/pkg/slog"
	pb "github.com/phil0522/taskd/proto"
	"google.golang.org/grpc"
	"gopkg.in/sevlyar/go-daemon.v0"
)

const (
	version    = "0.1"
	serverAddr = "127.0.0.1:6398"
)

func serve() {
	log.Printf("starting server")
	lis, err := net.Listen("tcp4", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer func() {
		log.Printf("server exit")
	}()
	grpcServer := grpc.NewServer()
	snipperServer := &server.SnippetServer{
		GrpcServer: grpcServer,
	}

	snipperServer.Initialize()
	pb.RegisterSnippetServiceServer(grpcServer, snipperServer)

	grpcServer.Serve(lis)
}

func formatSnippet(snippet *pb.ShellSnippet) {
	name := strings.ReplaceAll(snippet.GetSnippetName(), ":", "")
	desc := strings.ReplaceAll(snippet.GetSnippetDescription(), ":", "")
	cmd := snippet.GetSnippetCommand()
	//cmd := strings.ReplaceAll(snippet.GetSnippetCommand(), "\n", "")
	fmt.Printf("%s:%s:%s\n", name, desc, cmd)
}

func searchShell(ctx context.Context, client pb.SnippetServiceClient) {
	resp, err := client.SearchShellSnippet(ctx, &pb.ShellSnippetRequest{})

	//(ctx, &pb.ShellSnippetRequest{})
	if err != nil {
		log.Fatalf("failed to execute rpc %s", err.Error())
	}
	for _, snippet := range resp.ShellSnippets {
		formatSnippet(snippet)
	}
}

func killServer(ctx context.Context, client pb.SnippetServiceClient) {
	_, err := client.QuitServer(ctx, &pb.QuitServerRequest{})
	if err != nil {
		log.Fatalf("failed to quit server rpc %s", err.Error())
	}
}

func clientCall(callback func(context.Context, pb.SnippetServiceClient)) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create connection, %s", err.Error())
	}

	client := pb.NewSnippetServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	callback(ctx, client)
}

func executeCommand() {
	if flag.NArg() > 0 && flag.Arg(0) == "kill-server" {
		clientCall(killServer)
		return
	}
	clientCall(searchShell)
}

func main() {
	flag.BoolVar(&slog.VerboseLog, "verbose", false, "enable debugging log")
	flag.Parse()

	logger := slog.New("main")

	if flag.NArg() > 0 && flag.Arg(0) == "version" {
		fmt.Println(version)
		return
	}

	userRoot := os.Getenv("HOME")
	context := daemon.Context{
		PidFileName: filepath.Join(userRoot, "taskd.lock"),
		PidFilePerm: 0644,
		LogFileName: filepath.Join(userRoot, "taskd.log"),
		LogFilePerm: 0666,
		WorkDir:     userRoot,
	}

	child, _ := context.Search()
	if child != nil {
		logger.Debugf("Server has been already serving")
		executeCommand()
		return
	}

	if flag.NArg() > 0 && flag.Arg(0) == "kill-server" {
		log.Printf("server is not running, do nothing")
		return
	}

	child, err := context.Reborn()
	if err != nil {
		log.Fatal(err.Error())
	}

	if child != nil {
		time.Sleep(time.Second * 1)
		executeCommand()
	} else {
		serve()
		defer context.Release()
	}
}
