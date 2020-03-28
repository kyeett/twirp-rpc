package templateserver

import (
	"context"
	"log"

	pb "github.com/kyeett/twirp-rpc/examples/template/rpc/template"
	"go.uber.org/zap"
)

var _ pb.TemplateServicer = &Server{}

type Server struct {
	*zap.Logger
}

func New(logger *zap.Logger) *Server {
	return &Server{
		Logger: logger,
	}
}

func (s *Server) Update(context.Context, *pb.UpdateReq) (*pb.UpdateResp, error) {
	log.Fatal("Not implemented")
	return nil, nil
}
