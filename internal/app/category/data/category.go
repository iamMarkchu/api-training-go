package data

import (
	"api-training-go/api/common"
	"api-training-go/internal/app/category/biz"
	"api-training-go/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

func (r *categoryRepo) MGet(ctx context.Context, ids []uint64) (list []model.CategoryModel, err error) {
	list = make([]model.CategoryModel, 0, len(ids))
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

func (r *categoryRepo) List(ctx context.Context) (list []model.CategoryModel, total uint64, err error) {
	err = r.data.db.Where("status = ?", common.UserStatus_ACTIVE).Find(&list).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	var t int64
	err = r.data.db.Model(model.CategoryModel{}).Where("status = ?", common.UserStatus_ACTIVE).Count(&t).Error
	if err != nil {
		return
	}
	total = uint64(t)
	return
}

func (r *categoryRepo) GetOne(ctx context.Context, id uint64) (one model.CategoryModel, err error) {
	one = model.CategoryModel{Id: id, Status: uint8(common.UserStatus_ACTIVE)}
	if err = r.data.db.First(&one).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		r.log.WithContext(ctx).Error(err)
		return
	}
	return
}

func (r *categoryRepo) Create(ctx context.Context, m model.CategoryModel) (id uint64, err error) {
	if err = r.data.db.Create(&m).Error; err != nil {
		r.log.WithContext(ctx).Errorf("error:%s", err)
		return
	}
	id = m.Id
	return
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{data: data, log: log.NewHelper(logger)}
}
