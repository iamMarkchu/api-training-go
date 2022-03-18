package biz

import (
	pb_action "api-training-go/api/action"
	pb_category "api-training-go/api/category"
	pb_user "api-training-go/api/user"
	"context"
)

type CategoryRepo interface {
	MGetCategory(ctx context.Context, ids []uint64) (res map[uint64]*pb_category.CategoryModel, err error)
}

type UserRepo interface {
	MGetUser(ctx context.Context, uint64s []uint64) (res map[uint64]*pb_user.UserModel, err error)
}

type ActionRepo interface {
	MGetAction(ctx context.Context, ids []uint64) (res map[uint64]*pb_action.ActionModel, err error)
}
