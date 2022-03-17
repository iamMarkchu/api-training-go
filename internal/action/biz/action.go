package biz

import (
	pb "api-training-go/api/action"
	"api-training-go/api/common"
	"api-training-go/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type ActionUseCase struct {
	log    *log.Helper
	repo   ActionRepo
	ccRepo CategoryRepo
	ucRepo UserRepo
}

func (uc *ActionUseCase) CreateAction(ctx context.Context, req *pb.CreateActionRequest, uid uint64) (id uint64, err error) {
	m := model.ActionModel{
		CateId:      req.GetCateId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Status:      uint8(common.UserStatus_ACTIVE),
		UserId:      uid,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
	return uc.repo.Create(ctx, m)
}

func (uc *ActionUseCase) GetAction(ctx context.Context, req *pb.GetActionRequest) (data *pb.ActionModel, err error) {
	one, err := uc.repo.GetOne(ctx, req.GetId())
	if err != nil {
		return
	}
	ccMap, err := uc.ccRepo.MGetCategory(ctx, []uint64{one.CateId})
	ucMap, err := uc.ucRepo.MGetUser(ctx, []uint64{one.UserId})
	uc.log.Infof("ccMap:%v", ccMap)
	cate, _ := ccMap[one.CateId]
	user, _ := ucMap[one.UserId]
	data = &pb.ActionModel{
		Id:          one.Id,
		CateId:      one.CateId,
		Name:        one.Name,
		Description: one.Description,
		Status:      uint32(one.Status),
		User:        user,
		Ctime:       uint64(one.CreatedAt),
		Utime:       uint64(one.UpdatedAt),
		Category:    cate,
	}
	return
}

func NewActionUseCase(logger log.Logger, repo ActionRepo, ccRepo CategoryRepo, ucRepo UserRepo) *ActionUseCase {
	return &ActionUseCase{log: log.NewHelper(logger), repo: repo, ccRepo: ccRepo, ucRepo: ucRepo}
}

type ActionRepo interface {
	Create(ctx context.Context, m model.ActionModel) (uint64, error)
	GetOne(ctx context.Context, id uint64) (one model.ActionModel, err error)
}
