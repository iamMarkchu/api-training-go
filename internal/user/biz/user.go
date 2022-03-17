package biz

import (
	"api-training-go/api/common"
	pb "api-training-go/api/user"
	"api-training-go/internal/conf"
	"api-training-go/internal/model"
	"api-training-go/internal/pkg/auth/jwt"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"time"
)

type UserRepo interface {
	Create(ctx context.Context, m model.User) (id uint64, err error)
	GetUserByName(ctx context.Context, name string) (res model.User, err error)
	MGet(ctx context.Context, uids []uint64) (res []model.User, err error)
}

type UserUseCase struct {
	log  *log.Helper
	repo UserRepo
	conf *conf.Auth
}

func NewUserUseCase(logger log.Logger, repo UserRepo, conf *conf.Auth) *UserUseCase {
	return &UserUseCase{log: log.NewHelper(logger), repo: repo, conf: conf}
}

func (uc *UserUseCase) Register(ctx context.Context, req *pb.RegisterRequest) (id uint64, err error) {
	m := model.User{
		Name:      req.GetName(),
		Password:  req.GetPassword(),
		NickName:  req.GetNickName(),
		Status:    uint8(common.UserStatus_ACTIVE),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	return uc.repo.Create(ctx, m)
}

func (uc *UserUseCase) Login(ctx context.Context, req *pb.LoginRequest) (t string, ex int64, err error) {
	data, err := uc.repo.GetUserByName(ctx, req.GetName())
	if err != nil {
		return
	}
	if data.Id == 0 {
		err = common.ErrorUserNotFound("用户不存在")
		return
	}
	if data.Password != req.GetPassword() {
		err = common.ErrorUserNotFound("密码错误")
		return
	}
	ex = 86400
	claims := &jwt.ApiClaims{
		data.Id,
		jwtv4.RegisteredClaims{
			NotBefore: jwtv4.NewNumericDate(time.Now()),
			ExpiresAt: jwtv4.NewNumericDate(time.Now().Add(time.Second * 86400)),
			Issuer:    uc.conf.Key,
		},
	}
	token := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims)
	t, err = token.SignedString([]byte(uc.conf.Key))
	return
}

func (uc *UserUseCase) MGet(ctx context.Context, req *pb.MGetRequest) (res map[uint64]*pb.UserModel, err error) {
	list, err := uc.repo.MGet(ctx, req.GetUids())
	if err != nil {
		return
	}
	res = make(map[uint64]*pb.UserModel)
	for _, item := range list {
		res[item.Id] = &pb.UserModel{
			Id:   item.Id,
			Name: item.Name,
		}
	}
	return
}
