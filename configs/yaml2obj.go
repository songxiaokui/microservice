package configs

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
@Time    : 2021/2/8 16:04
@Author  : austsxk
@Email   : austsxk@163.com
@File    : yaml2obj.go
@Software: GoLand
*/

// 解析yaml文件为对象

func GetConfigPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	// 在cmd目录下编译，要获取该配置路径，就要从cmd目录开始计算文件所在路径
	return filepath.Join(filepath.Dir(filepath.Dir(pwd)), "configs", "config.yaml")
}

func GetYamlFile() ([]byte, error) {
	path := GetConfigPath()
	if path == "" {
		return []byte{}, errors.New("DB config error")
	}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, errors.New("DB config error")
	}
	return yamlFile, nil
}
