package data

import (
	"fmt"
	cf "microservice/configs"
	"testing"
)

/*
@Time    : 2021/2/8 16:47
@Author  : austsxk
@Email   : austsxk@163.com
@File    : user_biz_test.go
@Software: GoLand
*/

// 增加用户
func TestUserDaoImpl_Save(t *testing.T) {
	// 实例化对象
	UD := &UserDaoImpl{}
	cf.ManualSql()
	fmt.Println(cf.MysqlDB)
	// 创建用户
	var userArray []UserEntity
	userArray = append(userArray, UserEntity{Username: "宋晓奎", Email: "sxk@qq.com", Password: "sxk"})
	userArray = append(userArray, UserEntity{Username: "宋晓奎2", Email: "sxk@qq.com", Password: "sxk2"})
	for _, data := range userArray {
		err := UD.Save(&data)
		if err != nil {
			t.Errorf("mysql save error: %s", err)
			t.FailNow()
		}
		t.Log("new User ID is: ", data.ID)
	}
}
