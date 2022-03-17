package service

import (
	"api-training-go/api/common"
	"api-training-go/internal/category/biz"
	"api-training-go/internal/pkg/auth/jwt"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "api-training-go/api/category"
)

type CategoryService struct {
	pb.UnimplementedCategoryServer
	uc  *biz.CategoryUseCase
	log *log.Helper
}

func NewCategoryService(uc *biz.CategoryUseCase, logger log.Logger) *CategoryService {
	return &CategoryService{uc: uc, log: log.NewHelper(logger)}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryReply, error) {
	uid, ok := jwt.GetUserIdFromCtx(ctx)
	if !ok || uid <= 0 {
		return &pb.CreateCategoryReply{}, common.ErrorUserNotFound("未登录")
	}
	id, err := s.uc.CreateCategory(ctx, req, uid)
	if err != nil {
		return &pb.CreateCategoryReply{}, err
	}
	return &pb.CreateCategoryReply{
		Message: "ok",
		Data:    &common.ResponseData{Id: id},
	}, nil
}
func (s *CategoryService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryReply, error) {
	return &pb.UpdateCategoryReply{}, nil
}
func (s *CategoryService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryReply, error) {
	return &pb.DeleteCategoryReply{}, nil
}
func (s *CategoryService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryReply, error) {
	uid, ok := jwt.GetUserIdFromCtx(ctx)
	if !ok || uid <= 0 {
		return &pb.GetCategoryReply{}, common.ErrorUserNotFound("未登录")
	}
	data, err := s.uc.GetCategory(ctx, req, uid)
	if err != nil {
		return &pb.GetCategoryReply{}, err
	}
	return &pb.GetCategoryReply{
		Message: "ok",
		Data:    data,
	}, nil
}
func (s *CategoryService) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryReply, error) {
	list, total, err := s.uc.ListCategory(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListCategoryReply{
		Message: "ok",
		Data: &pb.ListCategoryReplyData{
			Total: total,
			List:  list,
		},
	}, nil
}
func (s *CategoryService) MGetCategory(ctx context.Context, req *pb.MGetCategoryRequest) (*pb.MGetCategoryReply, error) {
	data, err := s.uc.MGetCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.MGetCategoryReply{
		Message: "ok",
		Data:    data,
	}, nil
}
