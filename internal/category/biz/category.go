package biz

import (
	pb "api-training-go/api/category"
	"api-training-go/api/common"
	"api-training-go/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type CategoryUseCase struct {
	log  *log.Helper
	repo CategoryRepo
}

func (uc *CategoryUseCase) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest, uid uint64) (id uint64, err error) {
	uc.log.WithContext(ctx).Infof("params:%s, uid: %d", req, uid)
	m := model.CategoryModel{
		ParentId:    req.GetParentId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Status:      uint8(common.UserStatus_ACTIVE),
		UserId:      uid,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
	return uc.repo.Create(ctx, m)
}

func (uc *CategoryUseCase) GetCategory(ctx context.Context, req *pb.GetCategoryRequest, uid uint64) (data *pb.CategoryModel, err error) {
	uc.log.WithContext(ctx).Infof("params:%s, uid: %d", req, uid)
	one, err := uc.repo.GetOne(ctx, req.GetId())
	if err != nil {
		return
	}
	data = &pb.CategoryModel{
		Id:          one.Id,
		ParentId:    one.ParentId,
		Name:        one.Name,
		Description: one.Description,
		Status:      uint32(one.Status),
		UserId:      uint32(one.UserId),
		UserName:    "",
		Ctime:       uint64(one.CreatedAt),
		Utime:       uint64(one.UpdatedAt),
	}
	return
}

func (uc *CategoryUseCase) ListCategory(ctx context.Context) (list []*pb.CategoryModel, total uint64, err error) {
	uc.log.WithContext(ctx).Infof("params:%s")
	data, t, err := uc.repo.List(ctx)
	if err != nil {
		return
	}
	list = uc.transform(data)
	total = t
	return
}

func (uc *CategoryUseCase) transform(data []model.CategoryModel) (res []*pb.CategoryModel) {
	res = make([]*pb.CategoryModel, 0, len(data))
	for _, item := range data {
		res = append(res, &pb.CategoryModel{
			Id:          item.Id,
			ParentId:    item.ParentId,
			Name:        item.Name,
			Description: item.Description,
			Status:      uint32(item.Status),
			UserId:      uint32(item.UserId),
			UserName:    "",
			Ctime:       uint64(item.CreatedAt),
			Utime:       uint64(item.UpdatedAt),
		})
	}
	return
}

func (uc *CategoryUseCase) MGetCategory(ctx context.Context, req *pb.MGetCategoryRequest) (res map[uint64]*pb.CategoryModel, err error) {
	list, err := uc.repo.MGet(ctx, req.GetCateIds())
	if err != nil {
		return
	}
	res = make(map[uint64]*pb.CategoryModel)
	for _, item := range list {
		res[item.Id] = &pb.CategoryModel{
			Id:          item.Id,
			ParentId:    item.ParentId,
			Name:        item.Name,
			Description: item.Description,
			Status:      uint32(item.Status),
			UserId:      uint32(item.UserId),
			UserName:    "",
			Ctime:       uint64(item.CreatedAt),
			Utime:       uint64(item.UpdatedAt),
		}
	}
	return
}

func NewCategoryUseCase(logger log.Logger, repo CategoryRepo) *CategoryUseCase {
	return &CategoryUseCase{log: log.NewHelper(logger), repo: repo}
}

type CategoryRepo interface {
	Create(ctx context.Context, m model.CategoryModel) (id uint64, err error)
	GetOne(ctx context.Context, id uint64) (one model.CategoryModel, err error)
	List(ctx context.Context) (list []model.CategoryModel, total uint64, err error)
	MGet(ctx context.Context, ids []uint64) (list []model.CategoryModel, err error)
}
