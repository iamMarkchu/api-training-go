package service

import (
	"api-training-go/api/common"
	"api-training-go/internal/app/action/biz"
	"api-training-go/internal/pkg/auth/jwt"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "api-training-go/api/action"
)

type ActionService struct {
	pb.UnimplementedActionServer
	uc  *biz.ActionUseCase
	log *log.Helper
}

func NewActionService(uc *biz.ActionUseCase, logger log.Logger) *ActionService {
	return &ActionService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ActionService) CreateAction(ctx context.Context, req *pb.CreateActionRequest) (*pb.CreateActionReply, error) {
	uid, ok := jwt.GetUserIdFromCtx(ctx)
	if !ok || uid <= 0 {
		return &pb.CreateActionReply{}, common.ErrorUserNotFound("未登录")
	}
	id, err := s.uc.CreateAction(ctx, req, uid)
	if err != nil {
		return nil, err
	}
	return &pb.CreateActionReply{
		Message: "ok",
		Data:    &common.ResponseData{Id: id},
	}, nil
}
func (s *ActionService) UpdateAction(ctx context.Context, req *pb.UpdateActionRequest) (*pb.UpdateActionReply, error) {
	return &pb.UpdateActionReply{}, nil
}
func (s *ActionService) DeleteAction(ctx context.Context, req *pb.DeleteActionRequest) (*pb.DeleteActionReply, error) {
	return &pb.DeleteActionReply{}, nil
}
func (s *ActionService) GetAction(ctx context.Context, req *pb.GetActionRequest) (*pb.GetActionReply, error) {
	data, err := s.uc.GetAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.GetActionReply{
		Message: "ok",
		Data:    data,
	}, nil
}
func (s *ActionService) ListAction(ctx context.Context, req *pb.ListActionRequest) (*pb.ListActionReply, error) {
	return &pb.ListActionReply{}, nil
}

func (s *ActionService) MGetAction(ctx context.Context, req *pb.MGetActionRequest) (*pb.MGetActionReply, error) {
	data, err := s.uc.MGetAction(ctx, req)
	if err != nil {
		return &pb.MGetActionReply{}, err
	}
	return &pb.MGetActionReply{
		Message: "ok",
		Data:    data,
	}, nil
}
