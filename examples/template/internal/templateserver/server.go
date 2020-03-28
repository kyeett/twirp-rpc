package templateserver

import (
	"context"
	"log"

	pb "github.com/kyeett/twirp-rpc/examples/template/rpc/template"
)

var _ pb.TemplateServicer = &Server{}

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) Update(context.Context, *pb.UpdateReq) (*pb.UpdateResp, error) {
	log.Fatal("Not implemented")
	return nil, nil
}
