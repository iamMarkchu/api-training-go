package service

import (
	"api-training-go/internal/user/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "api-training-go/api/user"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, err := s.uc.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		Message: "ok",
	}, nil
}
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	t, ex, err := s.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Message: "ok",
		Data: &pb.LoginResponse_LoginData{
			Token:  t,
			Expire: uint64(ex),
		},
	}, nil
}

func (s *UserService) MGet(ctx context.Context, req *pb.MGetRequest) (*pb.MGetResponse, error) {
	data, err := s.uc.MGet(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.MGetResponse{
		Message: "ok",
		Data:    data,
	}, nil
}
