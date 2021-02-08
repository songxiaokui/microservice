package configs

/*
@Time    : 2021/2/8 15:56
@Author  : austsxk
@Email   : austsxk@163.com
@File    : mysql.go
@Software: GoLand
*/
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// 定义一个mysql对象，并且是一个连接池
var MysqlDB *gorm.DB

type Mysql struct {
	Host     string `yaml:"M_Host"`
	Password string `yaml:"M_Password"`
	UserName string `yaml:"M_User"`
	Port     string `yaml:"M_Port"`
}

func InitMysql(dbs string) (err error) {
	// 解析yaml
	yamlFile, err := GetYamlFile()
	if err != nil {
		return errors.Wrap(err, "config read error")
	}
	var m Mysql
	if err := yaml.Unmarshal(yamlFile, &m); err != nil {
		return errors.Wrap(err, "yaml unmarshal error")
	}
	// 创建连接mysql url
	mysqlUrl := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		m.UserName, m.Password, m.Host, m.Port, dbs)
	// 使用gorm创建连接对象
	MysqlDB, err = gorm.Open("mysql", mysqlUrl)
	if err != nil {
		return errors.Wrap(err, "mysql connect error")
	}
	MysqlDB.SingularTable(true)
	return
}

// 测试使用
func ManualSql() {
	MysqlDB, _ = gorm.Open("mysql", "root:mysql@(127.0.0.1:3306)/user?charset=utf8&parseTime=True&loc=Local")
	MysqlDB.SingularTable(true)
	return
}
