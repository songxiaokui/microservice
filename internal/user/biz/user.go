package biz

import (
	"time"
)

/*
@Time    : 2021/2/8 15:43
@Author  : austsxk
@Email   : austsxk@163.com
@File    : user.go
@Software: GoLand
*/

// model层
type UserEntity struct {
	ID         int64
	UserName   string
	Password   string
	Email      string
	CreateTime time.Time
}

// model对应数据库中的表名
func (UserEntity) TableName() string {
	return "user"
}

// 用户模型层对外提供新增和查寻的接口
type UserDaoInter interface {
	SelectByEmail(email string) (*UserEntity, error)
	Save(user *UserEntity) error
}

// 定义一个实现接口的类对象
type UserDaoImpl struct {
}

// 实现UserDaoInter接口
func (u *UserDaoImpl) SelectByEmail(email string) (*UserEntity, error) {
	user := &UserEntity{}
}
