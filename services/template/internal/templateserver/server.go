package templateserver

import (
	"context"

	pb "github.com/kyeett/twirp-rpc/services/template/rpc/template"
	"go.uber.org/zap"
)

var _ pb.TemplateServicer = &Server{}

type Server struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) Update(context.Context, *pb.UpdateReq) (*pb.UpdateResp, error) {
	s.logger.Fatal("Not implemented")
	return nil, nil
}
