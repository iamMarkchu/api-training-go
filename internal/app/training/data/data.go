package data

import (
	pb_action "api-training-go/api/action"
	pb_category "api-training-go/api/category"
	pb_user "api-training-go/api/user"
	"api-training-go/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewTrainingRepo, NewUserRepo, NewCategoryRepo, NewActionRepo)

// Data .
type Data struct {
	db *gorm.DB
	cc pb_category.CategoryClient
	uc pb_user.UserClient
	ac pb_action.ActionClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, clientConf *conf.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	db := openDB(c, logger)
	if db == nil {
		panic("db error")
	}
	return &Data{
		db: db,
		cc: NewCategoryGrpc(clientConf),
		uc: NewUserGrpc(clientConf),
		ac: NewActionGrpc(clientConf),
	}, cleanup, nil
}

func openDB(c *conf.Data, logger log.Logger) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Fatal("OPEN DB error")
		return nil
	}
	if c.Database.Debug {
		db = db.Debug()
	}
	return db
}
func NewCategoryGrpc(clientConf *conf.Client) pb_category.CategoryClient {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(clientConf.Grpc.Category.Addr))
	if err != nil {
		panic("grpc error")
	}
	return pb_category.NewCategoryClient(conn)
}

func NewUserGrpc(clientConf *conf.Client) pb_user.UserClient {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(clientConf.Grpc.User.Addr))
	if err != nil {
		panic("grpc error")
	}
	return pb_user.NewUserClient(conn)
}

func NewActionGrpc(clientConf *conf.Client) pb_action.ActionClient {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(clientConf.Grpc.Action.Addr))
	if err != nil {
		panic("grpc error")
	}
	return pb_action.NewActionClient(conn)
}
