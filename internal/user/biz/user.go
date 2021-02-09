package biz

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	dao "microservice/internal/user/data"
	"sync"
)

/*
@Time    : 2021/2/8 15:43
@Author  : austsxk
@Email   : austsxk@163.com
@File    : user.go
@Software: GoLand
*/
// 最终的依赖逻辑 server（服务抽象）--> service(服务组合) --> endpoint/biz(业务处理) --> data(数据持久化存储)
// 业务终端处理层,处理用户,依赖与DAO层，即data层

// 定义业务传输对象 DTO 对象
type UserInfoDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// 定义注册用户业务展示对象VO
type RegisterUserVO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// 定义通用的错误提示信息，该领域专用
var (
	ErrUserExists = errors.New("user is exists")
	ErrPassword   = errors.New("password is error")
)

var _ UserDoService = (*UserDoServiceImpl)(nil)

// 用户领域提供2个接口，就是用户登陆和用户注册
type UserDoService interface {
	// 用户登陆 携带请求上下午可以用来cancel 或 timeout，需要email和password 返回用户传输对象信息（不要直接返回数据库用户信息）
	Login(context.Context, string, string) (*UserInfoDTO, error)
	// 用户注册接口 从form表单重处理成视图对象，进行到后端处理
	Register(ctx context.Context, vo *RegisterUserVO) (*UserInfoDTO, error)
}

// 用户服务接口的实现层
type UserDoServiceImpl struct {
	// 服务层的实现，继承了用户数据层的实现,要用到领域层就是model的存储和数据持久
	daoUser dao.UserDaoInter
	// 注册时需要一个锁进行锁定
	mux sync.Mutex
}

// 将依赖的实现对象通过注入的方式，以便于拓展或mock测试
func NewUserDoServiceImpl(daoUser dao.UserDaoInter) *UserDoServiceImpl {
	return &UserDoServiceImpl{daoUser: daoUser}
}

func (u *UserDoServiceImpl) Login(ctx context.Context, email, password string) (*UserInfoDTO, error) {
	// 获取用户信息
	user, err := u.daoUser.SelectByEmail(email)
	// 如果获取成功，判断秘密是否正确,将DO对象转化为DTO对象，返回
	if err == nil {
		if user.Password == password {
			return &UserInfoDTO{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email}, nil
		}
		return nil, ErrPassword
	}
	// 如果不成功，则返回错误信息
	log.Printf("用户: %s, 不存在！", email)
	return nil, err

}

func (u *UserDoServiceImpl) Register(ctx context.Context, vo *RegisterUserVO) (*UserInfoDTO, error) {
	u.mux.Lock()
	defer u.mux.Unlock()
	// 先判断用户是否存在，如果存在，则直接返回，如果不存在，在进行用户存储返回

	userExists, err := u.daoUser.SelectByEmail(vo.Email)

	if err == nil && userExists == nil || err == gorm.ErrRecordNotFound {
		// 创建一个领域对象，model
		newUser := dao.UserEntity{
			Username: vo.Username,
			Password: vo.Password,
			Email:    vo.Email,
		}
		// save
		err := u.daoUser.Save(&newUser)
		if err == nil {
			return &UserInfoDTO{
				ID:       newUser.ID,
				Username: newUser.Username,
				Email:    newUser.Email,
			}, nil
		}
		return nil, err
	} else {
		err = ErrUserExists
	}

	return nil, err
}
