package data

import (
	pb_action "api-training-go/api/action"
	"api-training-go/internal/app/training/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type actionRepo struct {
	data *Data
	log  *log.Helper
}

func (r *categoryRepo) MGetAction(ctx context.Context, ids []uint64) (res map[uint64]*pb_action.ActionModel, err error) {
	resp, err := r.data.ac.MGetAction(ctx, &pb_action.MGetActionRequest{ActionIds: ids})
	if err != nil {
		r.log.WithContext(ctx).Error(err)
		return
	}
	res = resp.GetData()
	return
}

func NewActionRepo(data *Data, logger log.Logger) biz.ActionRepo {
	return &categoryRepo{data: data, log: log.NewHelper(logger)}
}
