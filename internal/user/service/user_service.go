package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	biz "microservice/internal/user/biz"
)

/*
@Time    : 2021/2/9 10:26
@Author  : austsxk
@Email   : austsxk@163.com
@File    : user_service.go
@Software: GoLand
*/

// 提供服务层，该层的服务是具体的业务实现，是业务领域的终端，将来给server提供服务，server只提供接口

type UserEndPoint struct {
	// 登陆的终端
	LoginEndPoint endpoint.Endpoint
	// 注册的终端
	RegisterEndPoint endpoint.Endpoint
}

// 定义在传输时，登陆时的请求和响应
type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	UserInfo *biz.UserInfoDTO `json:"user_info"`
}

// 彻底将业务进行解耦，不关心内部接口如何实现，只关心是否提供该接口
func MakeLoginEndPoint(userBiz biz.UserDoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (
		response interface{}, err error) {
		// 获取请求的参数
		req := request.(*LoginRequest)
		// 调用biz层业务处理逻辑
		userDTO, err := userBiz.Login(ctx, req.Email, req.Password)
		// 返回结果
		return &LoginResponse{userDTO}, err
	}
}

// 注册接口
type RegisterRequest struct {
	Email    string
	Password string
	Username string
}

type RegisterResponse struct {
	UserInfo *biz.UserInfoDTO `json:"user_info"`
}

func MakeRegisterEndPoint(userBiz biz.UserDoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (
		response interface{}, err error) {
		// 获取请求的参数
		req := request.(*RegisterRequest)
		// 调用biz层业务处理逻辑
		userDTO, err := userBiz.Register(ctx, &biz.RegisterUserVO{
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
		})
		// 返回结果
		return &RegisterResponse{userDTO}, err
	}
}
