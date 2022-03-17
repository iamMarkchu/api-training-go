package data

import (
	pb_user "api-training-go/api/user"
	"api-training-go/internal/action/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) MGetUser(ctx context.Context, uids []uint64) (res map[uint64]*pb_user.UserModel, err error) {
	resp, err := r.data.uc.MGet(ctx, &pb_user.MGetRequest{Uids: uids})
	if err != nil {
		return
	}
	res = resp.GetData()
	return
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{data: data, log: log.NewHelper(logger)}
}
