package configs

import (
	"fmt"
	"testing"
)

/*
@Time    : 2021/2/8 16:26
@Author  : austsxk
@Email   : austsxk@163.com
@File    : mysql_test.go
@Software: GoLand
*/

func TestGetConfigPath(t *testing.T) {
	path := GetConfigPath()
	fmt.Println(path)
}

func TestInitMysql(t *testing.T) {
	err := InitMysql("user")
	if err != nil {
		t.Errorf("db error: %s", err)
	}
}
