package server

import (
	"context"
	"log"

	pb "github.com/phil0522/taskd/proto"
	"google.golang.org/grpc"
)

// SnippetServer serve the snippet service.
type SnippetServer struct {
	GrpcServer *grpc.Server
	serveCount int32
}

// Serve starts the server until it is told to stop.
func (s *SnippetServer) Serve() {
}

// Stop stops the server
func (s *SnippetServer) Stop() {

}

// SearchShellSnippet returns all snippets match search criteria
func (s *SnippetServer) SearchShellSnippet(ctx context.Context, req *pb.ShellSnippetRequest) (*pb.ShellSnippetResponse, error) {
	resp := &pb.ShellSnippetResponse{}

	log.Printf("#%d, req: %v", s.serveCount, req)
	s.serveCount++

	snippet := &pb.ShellSnippet{
		SnippetName:        "SnippetName",
		SnippetCommand:     "echo Helloworld",
		SnippetDescription: "This is the test snippet",
	}
	resp.ShellSnippet = append(resp.ShellSnippet, snippet)
	if s.serveCount == 2 {
		s.GrpcServer.Stop()
	}
	return resp, nil
}
