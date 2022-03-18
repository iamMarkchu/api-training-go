package data

import (
	"api-training-go/api/common"
	"api-training-go/internal/app/user/biz"
	"api-training-go/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) GetUserByName(ctx context.Context, name string) (res model.User, err error) {
	res = model.User{}
	err = r.data.db.Where("name = ? and status = ?", name, common.UserStatus_ACTIVE).First(&res).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	return
}

func (r *userRepo) Create(ctx context.Context, m model.User) (id uint64, err error) {
	err = r.data.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("name = ?", m.Name).First(&m).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if m.Id > 0 {
			return common.ErrorUserExist("该用户已存在:%s", m.Name)
		}
		err = tx.Create(&m).Error
		if err != nil {
			return err
		}
		id = m.Id
		return nil
	})
	return
}

func (r *userRepo) MGet(ctx context.Context, uids []uint64) (list []model.User, err error) {
	list = make([]model.User, 0, len(uids))
	if len(uids) == 0 {
		return
	}
	err = r.data.db.Where("id in (?) and status = ?", uids, common.UserStatus_ACTIVE).Find(&list).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		r.log.WithContext(ctx).Error(err)
		return
	}
	return
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{data: data, log: log.NewHelper(logger)}
}
