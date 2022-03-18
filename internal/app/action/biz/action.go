package biz

import (
	pb "api-training-go/api/action"
	pb_category "api-training-go/api/category"
	"api-training-go/api/common"
	pb_user "api-training-go/api/user"
	"api-training-go/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"
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
	res := uc.transform(ctx, []model.ActionModel{one})
	data, _ = res[one.Id]
	return
}

func (uc *ActionUseCase) MGetAction(ctx context.Context, req *pb.MGetActionRequest) (res map[uint64]*pb.ActionModel, err error) {
	list, err := uc.repo.MGet(ctx, req.GetActionIds())
	if err != nil {
		return
	}
	res = uc.transform(ctx, list)
	return
}

func (uc *ActionUseCase) transform(ctx context.Context, list []model.ActionModel) (res map[uint64]*pb.ActionModel) {
	res = make(map[uint64]*pb.ActionModel, 0)
	if len(list) == 0 {
		return
	}
	cateIds := make([]uint64, 0)
	uids := make([]uint64, 0)
	for _, item := range list {
		cateIds = append(cateIds, item.CateId)
		uids = append(uids, item.UserId)
	}
	var (
		wg    errgroup.Group
		ccMap = make(map[uint64]*pb_category.CategoryModel)
		ucMap = make(map[uint64]*pb_user.UserModel)
	)
	wg.Go(func() (err error) {
		ccMap, err = uc.ccRepo.MGetCategory(ctx, cateIds)
		return
	})
	wg.Go(func() (err error) {
		ucMap, err = uc.ucRepo.MGetUser(ctx, uids)
		return
	})
	wg.Wait()
	for _, item := range list {
		user, _ := ucMap[item.UserId]
		cate, _ := ccMap[item.CateId]
		res[item.Id] = &pb.ActionModel{
			Id:          item.Id,
			CateId:      item.CateId,
			Name:        item.Name,
			Description: item.Description,
			Status:      uint32(item.Status),
			User:        user,
			Ctime:       uint64(item.CreatedAt),
			Utime:       uint64(item.UpdatedAt),
			Category:    cate,
		}
	}
	return
}

func NewActionUseCase(logger log.Logger, repo ActionRepo, ccRepo CategoryRepo, ucRepo UserRepo) *ActionUseCase {
	return &ActionUseCase{log: log.NewHelper(logger), repo: repo, ccRepo: ccRepo, ucRepo: ucRepo}
}

type ActionRepo interface {
	Create(ctx context.Context, m model.ActionModel) (uint64, error)
	GetOne(ctx context.Context, id uint64) (one model.ActionModel, err error)
	MGet(ctx context.Context, ids []uint64) (list []model.ActionModel, err error)
}
