package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"mxshop_srvs/user_srv/global"
	models "mxshop_srvs/user_srv/model"
	"mxshop_srvs/user_srv/proto"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

type UserServer struct {
	proto.UnimplementedUserServer
}

// 将model转换成可返回的响应类型值
func ModeltoResponse(user models.User) proto.UserInfoResponse {
	//grpc的message中字段有默认值，不能随便赋值nil进去，这里要搞清楚哪些字段有默认值
	UserInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		//Birthday: uint64(user.Birthday), //
	}
	if user.Birthday != nil {
		UserInfoRsp.Birthday = uint64(user.Birthday.Unix())

	}
	return UserInfoRsp
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	//获取用户列表
	//查询用户列表
	//获取全局数据库连接
	var users []models.User
	result := global.DB.Find(&users)
	if result.Error != nil { //出现报错
		return nil, result.Error
	}
	fmt.Println("用户列表")
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	//分页,取用户数据，参考gorm官方文档
	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	for _, user := range users {
		userInfoRsp := ModeltoResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}

func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	//通过用户ID查询用户
	var user models.User
	result := global.DB.Where(&models.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在") //grpc状态码
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoRsp := ModeltoResponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user models.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在") //grpc状态码
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoRsp := ModeltoResponse(user)
	return &userInfoRsp, nil

}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	//新建用户
	//1.首先查询是否存在，通过手机号来查询用户是否存在
	var user models.User
	res := global.DB.Where(&models.User{Mobile: req.Mobile}).First(user)
	if res.RowsAffected == 1 {
		//用户已存在
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	//没有查询到
	user.Mobile = req.Mobile
	user.Nickname = req.Nickname

	//密码加密，盐值加密，将用户密码变成随机数加用户密码
	options := &password.Options{16, 100, 32, sha512.New} //sha512为更安全的做法
	salt, encodedPwd := password.Encode(req.Password, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	res = global.DB.Create(&user) //创建用户
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}

	userinfoRsp := ModeltoResponse(user)
	return &userinfoRsp, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	//更新用户信息
	var user models.User
	//首先查询用户是否存在
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在") //grpc状态码
	}
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	//更新信息
	Birthday := time.Unix(int64(req.Birthday), 0) //int转time
	user.Nickname = req.Nickname
	user.Birthday = &Birthday
	user.Gender = req.Gender
	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil

}

// 检查用户密码
func (s *UserServer) CheckUserPassword(ctx context.Context, req *proto.PasswordCheck) (*proto.CheckResponse, error) {
	//校验密码
	options := &password.Options{10, 100, 32, sha512.New}
	passwordinfo := strings.Split(req.Encryptedpassword, "$")
	check := password.Verify(req.Password, passwordinfo[2], passwordinfo[3], options)
	return &proto.CheckResponse{Success: check}, nil
}

//func (s *UserServer) mustEmbedUnimplementedUserServer(){
//
//}
