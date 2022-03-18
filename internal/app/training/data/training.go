package data

import (
	"api-training-go/api/common"
	"api-training-go/internal/app/training/biz"
	"api-training-go/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type trainingRepo struct {
	data *Data
	log  *log.Helper
}

func (r *trainingRepo) GetListByUser(ctx context.Context, uid uint64) (list []model.TrainingModel, total uint64, err error) {
	err = r.data.db.Model(model.TrainingModel{}).Where("user_id = ?", uid).Find(&list).Error
	if err != nil {
		r.log.WithContext(ctx).Error(err)
	}
	var t int64
	err = r.data.db.Model(model.TrainingModel{}).Where("user_id = ?", uid).Count(&t).Error
	if err != nil {
		r.log.WithContext(ctx).Error(err)
	}
	total = uint64(t)
	return
}

func (r *trainingRepo) GetItemList(ctx context.Context, ids []uint64) (list []model.TrainingItemModel, err error) {
	err = r.data.db.Model(model.TrainingItemModel{}).Where("training_id in (?)", ids).Find(&list).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		r.log.WithContext(ctx).Errorf("GetItemList error:%s, ids:%d", err, ids)
		return
	}
	return
}

func (r *trainingRepo) Delete(ctx context.Context, id uint64) (err error) {
	err = r.data.db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&model.TrainingModel{}).Where("id = ?", id).Updates(map[string]interface{}{
			"status":     common.UserStatus_INACTIVE,
			"updated_at": time.Now().Unix(),
		}).Error
		if err != nil {
			r.log.WithContext(ctx).Errorf("Delete error:%s, id:%d", err, id)
			return
		}
		err = tx.Model(&model.TrainingItemModel{}).Where("training_id = ?", id).Updates(map[string]interface{}{
			"status":     common.UserStatus_INACTIVE,
			"updated_at": time.Now().Unix(),
		}).Error
		if err != nil {
			r.log.WithContext(ctx).Errorf("Delete[2] error:%s, id:%d", err, id)
			return
		}
		return nil
	})
	return
}

func (r *trainingRepo) GetOne(ctx context.Context, id uint64) (one model.TrainingModel, err error) {
	one.Id = id
	if err = r.data.db.First(&one).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		r.log.WithContext(ctx).Errorf("GetOne Error:%s, id: %d", err, id)
		return
	}
	return
}

func (r *trainingRepo) Create(ctx context.Context, m model.TrainingModel, mi []model.TrainingItemModel) (id uint64, err error) {
	err = r.data.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&m).Error; err != nil {
			r.log.WithContext(ctx).Errorf("err: %s", err)
			return err
		}
		for k := range mi {
			mi[k].TrainingId = m.Id
		}
		if err = tx.CreateInBatches(&mi, len(mi)).Error; err != nil {
			r.log.WithContext(ctx).Errorf("err: %s", err)
			return err
		}
		id = m.Id
		return nil
	})
	return
}

func NewTrainingRepo(data *Data, logger log.Logger) biz.TrainingRepo {
	return &trainingRepo{data: data, log: log.NewHelper(logger)}
}
