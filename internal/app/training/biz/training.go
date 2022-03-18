package biz

import (
	"api-training-go/api/common"
	"api-training-go/api/training"
	"api-training-go/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type TrainingRepo interface {
	Create(ctx context.Context, m model.TrainingModel, mi []model.TrainingItemModel) (id uint64, err error)
	GetOne(ctx context.Context, id uint64) (one model.TrainingModel, err error)
	Delete(ctx context.Context, id uint64) (err error)
	GetItemList(ctx context.Context, ids []uint64) (list []model.TrainingItemModel, err error)
	GetListByUser(ctx context.Context, uid uint64) (list []model.TrainingModel, total uint64, err error)
}

type TrainingUseCase struct {
	log    *log.Helper
	repo   TrainingRepo
	ucRepo UserRepo
	acRepo ActionRepo
}

func (uc *TrainingUseCase) CreateTraining(ctx context.Context, req *training.CreateTrainingRequest, uid uint64) (id uint64, err error) {
	uc.log.WithContext(ctx).Infof("CreateTraining: %s", req)
	m := model.TrainingModel{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		UserId:      uid,
		Status:      uint8(common.UserStatus_ACTIVE),
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
	mi := make([]model.TrainingItemModel, 0, len(req.GetList()))
	for _, item := range req.GetList() {
		mi = append(mi, model.TrainingItemModel{
			GroupId:    item.GetGroupId(),
			TrainingId: item.GetTrainingId(),
			ActionId:   item.GetActionId(),
			CountType:  uint8(item.GetCountType()),
			Weight:     item.GetWeight(),
			CountNum:   item.GetCountNum(),
			UserId:     uid,
			Status:     uint8(common.UserStatus_ACTIVE),
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		})
	}
	return uc.repo.Create(ctx, m, mi)
}

func (uc *TrainingUseCase) DeleteTraining(ctx context.Context, req *training.DeleteTrainingRequest, uid uint64) (id uint64, err error) {
	uc.log.WithContext(ctx).Infof("DeleteTraining params: %s, uid: %d", req, uid)
	// 获取训练数据
	one, err := uc.repo.GetOne(ctx, req.GetId())
	if err != nil {
		return
	}
	if one.UserId != uid {
		err = errors.New(500, "不允许修改", "不允许修改")
		return
	}
	err = uc.repo.Delete(ctx, req.GetId())
	if err != nil {
		return
	}
	id = one.Id
	return
}

func (uc *TrainingUseCase) GetTraining(ctx context.Context, req *training.GetTrainingRequest, uid uint64) (res *training.TrainingModel, err error) {
	uc.log.WithContext(ctx).Infof("GetTraining params: %s, uid:%d", req, uid)
	if req.GetId() <= 0 {
		return
	}
	one, err := uc.repo.GetOne(ctx, req.GetId())
	if err != nil {
		return
	}
	itemList, err := uc.repo.GetItemList(ctx, []uint64{req.GetId()})
	if err != nil {
		return
	}
	resMap := uc.transform(ctx, []model.TrainingModel{one}, itemList)
	res, _ = resMap[one.Id]
	return
}

func (uc *TrainingUseCase) ListTraining(ctx context.Context, req *training.ListTrainingRequest, uid uint64) (res []*training.TrainingModel, total uint64, err error) {
	uc.log.WithContext(ctx).Infof("ListTraining params: %s, uid:%d", req, uid)
	res = make([]*training.TrainingModel, 0)
	list, total, err := uc.repo.GetListByUser(ctx, uid)
	if err != nil {
		return
	}
	trainingIds := make([]uint64, 0)
	for _, item := range list {
		trainingIds = append(trainingIds, item.Id)
	}
	itemList, err := uc.repo.GetItemList(ctx, trainingIds)
	if err != nil {
		return
	}
	resMap := uc.transform(ctx, list, itemList)
	for _, item := range list {
		tr, _ := resMap[item.Id]
		res = append(res, tr)
	}
	return
}

func (uc *TrainingUseCase) transform(ctx context.Context, list []model.TrainingModel, itemList []model.TrainingItemModel) (res map[uint64]*training.TrainingModel) {
	res = make(map[uint64]*training.TrainingModel)
	// 获取用户
	uids := make([]uint64, 0)
	for _, item := range list {
		uids = append(uids, item.UserId)
	}
	ucMap, _ := uc.ucRepo.MGetUser(ctx, uids)
	// 获取动作
	actionIds := make([]uint64, 0)
	for _, item := range itemList {
		actionIds = append(actionIds, item.ActionId)
	}
	acMap, _ := uc.acRepo.MGetAction(ctx, actionIds)
	for _, item := range list {
		user, _ := ucMap[item.UserId]
		res[item.Id] = &training.TrainingModel{
			Id:          item.Id,
			Title:       item.Title,
			Description: item.Description,
			User:        user,
		}
	}
	for _, item := range itemList {
		action, _ := acMap[item.ActionId]
		if _, ok := res[item.TrainingId]; !ok {
			continue
		}
		if res[item.TrainingId].List == nil {
			res[item.TrainingId].List = make([]*training.TrainingItem, 0)
		}
		res[item.TrainingId].List = append(res[item.TrainingId].List, &training.TrainingItem{
			Id:         item.Id,
			GroupId:    item.GroupId,
			TrainingId: item.TrainingId,
			ActionId:   item.ActionId,
			CountType:  uint32(item.CountType),
			Weight:     item.Weight,
			CountNum:   item.CountNum,
			Action:     action,
		})
	}
	return res
}

func NewTrainingUseCase(repo TrainingRepo, logger log.Logger, ucRepo UserRepo, acRepo ActionRepo) *TrainingUseCase {
	return &TrainingUseCase{
		repo:   repo,
		log:    log.NewHelper(logger),
		ucRepo: ucRepo,
		acRepo: acRepo,
	}
}
