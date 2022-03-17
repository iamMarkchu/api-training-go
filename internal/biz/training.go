package biz

import (
	"api-training-go/api/common"
	"api-training-go/api/training"
	"api-training-go/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type TrainingRepo interface {
	Create(ctx context.Context, m model.TrainingModel, mi []model.TrainingItemModel) (id uint64, err error)
}

type TrainingUseCase struct {
	log  *log.Helper
	repo TrainingRepo
}

func (c *TrainingUseCase) CreateTraining(ctx context.Context, req *training.CreateTrainingRequest, uid uint64) (id uint64, err error) {
	c.log.WithContext(ctx).Infof("CreateTraining: %s", req)
	m := model.TrainingModel{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		UserId:      uid,
		Status:      uint8(common.UserStatus_ACTIVE),
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
	mi := make([]model.TrainingItemModel, 0, len(req.GetList()))
	for _, item := range req.GetList() {
		mi = append(mi, model.TrainingItemModel{
			GroupId:    item.GetGroupId(),
			TrainingId: item.GetTrainingId(),
			ActionId:   item.GetActionId(),
			CountType:  uint8(item.GetCountType()),
			Weight:     item.GetWeight(),
			CountNum:   item.GetCountNum(),
			UserId:     0,
			Status:     uint8(common.UserStatus_ACTIVE),
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		})
	}
	return c.repo.Create(ctx, m, mi)
}

func NewTrainingUseCase(repo TrainingRepo, logger log.Logger) *TrainingUseCase {
	return &TrainingUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
