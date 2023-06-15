package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/omalloc/kratos-layout/api/helloworld/v1"
	"github.com/omalloc/kratos-layout/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	log *log.Helper
	uc  *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(logger log.Logger, uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{
		log: log.NewHelper(logger),
		uc:  uc,
	}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

// Check greeter health checker, implement health.Checker interface.
func (s *GreeterService) Check(ctx context.Context) error {
	s.log.Info("check greeter health")
	return nil
}
