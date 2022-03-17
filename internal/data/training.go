package data

import (
	"api-training-go/internal/biz"
	"api-training-go/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type trainingRepo struct {
	data *Data
	log  *log.Helper
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
