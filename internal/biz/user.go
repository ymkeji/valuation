package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id         uint64    ` json:"id,omitempty"`
	UserName   string    ` json:"username,omitempty"`
	Password   string    ` json:"password,omitempty"`
	Avatar     string    ` json:"avatar,omitempty"`
	Expiration time.Time ` json:"expiration,omitempty"`
	CreateTime time.Time ` json:"createTime,omitempty"`
}
type UserRepo interface {
	FindByName(context.Context, string) (*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase  new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) FindByName(ctx context.Context, username string) (*User, error) {
	uc.log.WithContext(ctx).Infof("username: %s", username)
	return uc.repo.FindByName(ctx, username)
}
