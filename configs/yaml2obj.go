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
	return filepath.Join(filepath.Dir(pwd), "configs", "config.yaml")
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
