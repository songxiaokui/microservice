package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	us "microservice/internal/user/service"
	"net/http"
	"os"
)

/*
@Time    : 2021/2/9 09:37
@Author  : austsxk
@Email   : austsxk@163.com
@File    : user.go
@Software: GoLand
*/

var (
	ErrBadRequest = errors.New("invalid request parameter")
)

// 自定义http请求的处理器 mux
func MakeHttpHandler(ctx context.Context, endpoint *us.UserEndPoint) http.Handler {
	// 创建多路复用的路由
	route := mux.NewRouter()
	// 日志
	kitLog := log.NewLogfmtLogger(os.Stderr)
	kitLog = log.With(kitLog, "ts", log.DefaultTimestampUTC)
	kitLog = log.With(kitLog, "caller", log.DefaultCaller)
	// 添加kit http服务的可选参数
	opt := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitLog)),
		kithttp.ServerErrorEncoder(encodeError),
	}
	// 添加处理请求,路由，已经处理的kit http服务，服务包括终端处理，服务的请求，服务的响应已经其他可选参数
	route.Methods("POST").Path("/login").Handler(kithttp.NewServer(
		endpoint.LoginEndPoint,
		encodeLoginRequest,
		encodeJSONResponse,
		opt...,
	))
	// 注册
	route.Methods("POST").Path("/register").Handler(kithttp.NewServer(
		endpoint.RegisterEndPoint,
		encodeRegisterRequest,
		encodeJSONResponse,
		opt...,
	))
	return route
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	// 处理成json
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func encodeLoginRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	// 从form表单中获取参数
	username := req.FormValue("username")
	password := req.FormValue("password")

	// 组装成终端处理的登陆的请求对象
	if username == "" || password == "" {
		return nil, ErrBadRequest
	}
	// 组装成终端处理响应的响应对象
	return &us.LoginRequest{Password: password,
		Email: username,
	}, nil
}

// 处理注册
func encodeRegisterRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	email := req.FormValue("email")
	if username == "" || password == "" || email == "" {
		return nil, ErrBadRequest
	}
	return &us.RegisterRequest{Username: username,
		Password: password,
		Email:    email,
	}, nil
}

func encodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
