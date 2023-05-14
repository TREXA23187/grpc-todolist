package handler

import (
	"github.com/gin-gonic/gin"
	"grpc-todolist/user/internal/repository"
	"grpc-todolist/user/internal/service"
	"grpc-todolist/user/pkg/costum_error"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) UserLogin(ctx *gin.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = costum_error.Success

	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = costum_error.Error
		return resp, err
	}

	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

func (*UserService) UserRegister(ctx *gin.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = costum_error.Success

	err = user.CreateUser(req)
	if err != nil {
		resp.Code = costum_error.Error
		return resp, err
	}

	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}
