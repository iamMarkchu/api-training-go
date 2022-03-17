package data

import (
	pb_category "api-training-go/api/category"
	"api-training-go/internal/action/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

func (r *categoryRepo) MGetCategory(ctx context.Context, ids []uint64) (res map[uint64]*pb_category.CategoryModel, err error) {
	resp, err := r.data.cc.MGetCategory(ctx, &pb_category.MGetCategoryRequest{CateIds: ids})
	if err != nil {
		return
	}
	res = resp.GetData()
	return
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{data: data, log: log.NewHelper(logger)}
}
