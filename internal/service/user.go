package service

import (
	"context"
	"strconv"

	pb "valuation/api/valuation/v1"
	"valuation/internal/biz"
	"valuation/pkg/convertx"
	"valuation/pkg/errorx"
	"valuation/pkg/jwt"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	uc  *biz.UserUsecase
	log *log.Helper
	pb.UnimplementedUserServer
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginReply, error) {
	user, err := s.uc.FindByName(ctx, req.Username)
	errorx.Dangerous(err)
	if user == nil {
		errorx.Bomb(201, "User not exist")
	}

	if err := bcrypt.CompareHashAndPassword(convertx.String2Bytes(user.Password), convertx.String2Bytes(req.Password)); err != nil {
		errorx.Bomb(201, "Password mismatch")
	}

	token, err := jwt.EncodeToken(strconv.FormatUint(user.Id, 10))
	errorx.Dangerous(err)

	return &pb.UserLoginReply{
		AccessToken: token,
	}, nil
}
