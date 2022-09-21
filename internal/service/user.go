package service

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"valuation/pkg/errorx"

	pb "valuation/api/valuation/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginReply, error) {
	mySigningKey := []byte("AllYourBase")
	o := &struct {
		signingMethod jwt.SigningMethod
		claims        jwt.Claims
	}{
		signingMethod: jwt.SigningMethodHS256,
		claims: jwt.RegisteredClaims{
			ID: "1",
		},
	}
	token := jwt.NewWithClaims(o.signingMethod, o.claims)
	tokenString, err := token.SignedString(mySigningKey)
	errorx.Dangerous(err)
	return &pb.UserLoginReply{
		AccessToken: tokenString,
	}, nil
}
