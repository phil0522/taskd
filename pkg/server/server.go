package server

import (
	"context"
	"time"

	"github.com/phil0522/taskd/pkg/slog"
	pb "github.com/phil0522/taskd/proto"
	"google.golang.org/grpc"
)

// SnippetServer serve the snippet service.
type SnippetServer struct {
	GrpcServer    *grpc.Server
	serveCount    int32
	snipperManger *snipperManger
}

var (
	logger = slog.New("server")
)

// Initialize initialize the server
func (s *SnippetServer) Initialize() {
	logger.Debugf("initialize snippet manger.")
	s.snipperManger = &snipperManger{}
	s.snipperManger.initialize()
}

// SearchShellSnippet returns all snippets match search criteria
func (s *SnippetServer) SearchShellSnippet(ctx context.Context, req *pb.ShellSnippetRequest) (*pb.ShellSnippetResponse, error) {
	resp := &pb.ShellSnippetResponse{}

	logger.Debugf("#%d, req: %v", s.serveCount, req)
	s.serveCount++

	for _, snippet := range s.snipperManger.snippetsByCategory[req.Category] {
		logger.Debugf("add snippet %v", snippet)
		resp.ShellSnippets = append(resp.ShellSnippets, snippet)
	}

	return resp, nil
}

// QuitServer quits the server.
func (s *SnippetServer) QuitServer(ctx context.Context, in *pb.QuitServerRequest) (*pb.QuitServerResponse, error) {
	go func() {
		time.Sleep(time.Second * 1)
		s.GrpcServer.Stop()
	}()
	return &pb.QuitServerResponse{}, nil
}
