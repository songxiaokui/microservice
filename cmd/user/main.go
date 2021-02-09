package main

/*
@Time    : 2021/2/8 14:57
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

import (
	"context"
	"fmt"
	"log"
	"microservice/configs"
	"microservice/internal/user/biz"
	"microservice/internal/user/data"
	"microservice/internal/user/server"
	"microservice/internal/user/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)
	// 初始化数据库
	err := configs.InitMysql("user")
	if err != nil {
		log.Fatal("db init error:", err.Error())
	}
	// dao层的实现是一个dao数据存储层的接口
	bizServerImpl := biz.NewUserDoServiceImpl(&data.UserDaoImpl{})

	// BFF层
	serviceServerImpl := service.UserEndPoint{
		LoginEndPoint:    service.MakeLoginEndPoint(bizServerImpl),
		RegisterEndPoint: service.MakeRegisterEndPoint(bizServerImpl),
	}
	// 服务层
	serverHandler := server.MakeHttpHandler(ctx, &serviceServerImpl)

	// 启动服务并使用chan进行组册主进程
	go func() {
		fmt.Println("micro server of user is running at 127.0.0.1:8888...")
		errChan <- http.ListenAndServe("127.0.0.1:8888", serverHandler)
	}()
	// 监听信号事件
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// 持续阻塞，一旦有错误或者退出信号，则服务退出
	<-errChan
	fmt.Println("server will shutdown...")
}
