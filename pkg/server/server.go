package server

import (
	"context"
	"log"
	"time"

	pb "github.com/phil0522/taskd/proto"
	"google.golang.org/grpc"
)

// SnippetServer serve the snippet service.
type SnippetServer struct {
	GrpcServer *grpc.Server
	serveCount int32

	snipperManger *snipperManger
}

// Initialize initialize the server
func (s *SnippetServer) Initialize() {
	log.Printf("initialize snippet manger.")
	s.snipperManger = &snipperManger{}
	s.snipperManger.initialize()
}

// SearchShellSnippet returns all snippets match search criteria
func (s *SnippetServer) SearchShellSnippet(ctx context.Context, req *pb.ShellSnippetRequest) (*pb.ShellSnippetResponse, error) {
	resp := &pb.ShellSnippetResponse{}

	log.Printf("#%d, req: %v", s.serveCount, req)
	s.serveCount++

	for _, snippet := range s.snipperManger.globalShellSnippets {
		log.Printf("add snippet %v", snippet)
		resp.ShellSnippet = append(resp.ShellSnippet, snippet)
	}
	go func() {
		time.Sleep(time.Second * 1)
		s.GrpcServer.Stop()
	}()

	return resp, nil
}
