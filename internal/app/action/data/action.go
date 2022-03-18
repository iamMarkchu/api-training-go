package data

import (
	"api-training-go/api/common"
	"api-training-go/internal/app/action/biz"
	"api-training-go/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type actionRepo struct {
	data *Data
	log  *log.Helper
}

func (r *actionRepo) MGet(ctx context.Context, ids []uint64) (list []model.ActionModel, err error) {
	list = make([]model.ActionModel, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	err = r.data.db.Where("id in (?) and status = ?", ids, common.UserStatus_ACTIVE).Find(&list).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		r.log.WithContext(ctx).Error(err)
		return
	}
	return
}

func (r *actionRepo) GetOne(ctx context.Context, id uint64) (one model.ActionModel, err error) {
	one = model.ActionModel{Id: id, Status: uint8(common.UserStatus_ACTIVE)}
	err = r.data.db.First(&one).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		r.log.WithContext(ctx).Error(err)
		return
	}
	return
}

func NewActionRepo(data *Data, logger log.Logger) biz.ActionRepo {
	return &actionRepo{data: data, log: log.NewHelper(logger)}
}

func (r *actionRepo) Create(ctx context.Context, m model.ActionModel) (id uint64, err error) {
	err = r.data.db.Create(&m).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		r.log.WithContext(ctx).Error(err)
		return
	}
	id = m.Id
	return
}
