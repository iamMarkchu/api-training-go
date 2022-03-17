package service

import (
	"api-training-go/api/common"
	"api-training-go/internal/biz"
	"api-training-go/internal/pkg/auth/jwt"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "api-training-go/api/training"
)

type TrainingService struct {
	pb.UnimplementedTrainingServer
	log *log.Helper
	uc  *biz.TrainingUseCase
}

func NewTrainingService(uc *biz.TrainingUseCase, logger log.Logger) *TrainingService {
	return &TrainingService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *TrainingService) CreateTraining(ctx context.Context, req *pb.CreateTrainingRequest) (*pb.CreateTrainingReply, error) {
	uid, ok := jwt.GetUserIdFromCtx(ctx)
	if !ok || uid <= 0 {
		return &pb.CreateTrainingReply{}, common.ErrorUserNotFound("未登录")
	}
	id, err := s.uc.CreateTraining(ctx, req, uid)
	if err != nil {
		return &pb.CreateTrainingReply{}, err
	}
	return &pb.CreateTrainingReply{
		Message: "ok",
		Data:    &common.ResponseData{Id: id},
	}, nil
}
func (s *TrainingService) UpdateTraining(ctx context.Context, req *pb.UpdateTrainingRequest) (*pb.UpdateTrainingReply, error) {
	return &pb.UpdateTrainingReply{}, nil
}
func (s *TrainingService) DeleteTraining(ctx context.Context, req *pb.DeleteTrainingRequest) (*pb.DeleteTrainingReply, error) {
	return &pb.DeleteTrainingReply{}, nil
}
func (s *TrainingService) GetTraining(ctx context.Context, req *pb.GetTrainingRequest) (*pb.GetTrainingReply, error) {
	return &pb.GetTrainingReply{}, nil
}
func (s *TrainingService) ListTraining(ctx context.Context, req *pb.ListTrainingRequest) (*pb.ListTrainingReply, error) {
	return &pb.ListTrainingReply{}, nil
}
