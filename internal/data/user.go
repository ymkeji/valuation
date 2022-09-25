package data

import (
	"context"
	"time"

	"valuation/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type User struct {
	Id         uint64    ` json:"id,omitempty"`
	UserName   string    ` json:"username,omitempty"`
	Password   string    ` json:"password,omitempty"`
	Avatar     string    ` json:"avatar,omitempty"`
	Expiration time.Time ` json:"expiration,omitempty"`
	CreateTime time.Time ` json:"createTime,omitempty"`
}

func (u *userRepo) FindByName(ctx context.Context, username string) (*biz.User, error) {
	var user User
	result := u.data.db.Where("username = ?", username).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &biz.User{
		Id:         user.Id,
		UserName:   user.UserName,
		Password:   user.Password,
		Avatar:     user.Avatar,
		Expiration: user.Expiration,
		CreateTime: user.CreateTime,
	}, nil
}
