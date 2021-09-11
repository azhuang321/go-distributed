package handler

import (
	"context"
	"go.uber.org/zap"
	"user_srv/library"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"user_srv/database"
	"user_srv/model"
	"user_srv/proto"
)

type UserService struct{}

func convertModelUserToResponseUser(user model.User) *proto.UserInfoResponse {
	respUser := proto.UserInfoResponse{}
	respUser.Id = int32(user.ID)
	respUser.Password = user.Password
	respUser.Mobile = user.Mobile
	respUser.Role = uint32(user.Role)
	respUser.Nickname = user.Nickname
	respUser.Gender = uint32(user.Gender)
	respUser.Birthday = uint64(user.Birthday)
	return &respUser
}

func (u *UserService) GetUserByMobile(ctx context.Context, request *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	db, err := database.GetDB()
	if err != nil {
		zap.S().Errorf("数据库错误:%s\n", err.Error())
		return nil, status.Error(codes.DataLoss, "获取数据库出错")
	}
	db.Where("mobile = ?", request.Mobile).First(&user)
	if user.ID < 1 {
		return nil, status.Error(codes.NotFound, "未找到用户")
	}
	userInfoResp := convertModelUserToResponseUser(user)
	return userInfoResp, nil
}

func (u *UserService) GetUserById(ctx context.Context, request *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	db, err := database.GetDB()
	if err != nil {
		zap.S().Errorf("数据库错误:%s\n", err.Error())
		return nil, status.Error(codes.DataLoss, "获取数据库出错")
	}
	db.First(&user, request.Id)
	if user.ID < 1 {
		return nil, status.Error(codes.NotFound, "未找到用户")
	}
	userInfoResp := convertModelUserToResponseUser(user)
	return userInfoResp, nil
}

func (u *UserService) CreateUser(ctx context.Context, info *proto.CreateUserInfoRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	db, err := database.GetDB()
	if err != nil {
		zap.S().Errorf("数据库错误:%s\n", err.Error())
		return nil, status.Error(codes.DataLoss, "获取数据库出错")
	}
	db.Where("mobile = ?", info.Mobile).First(&user)
	if user.ID > 0 {
		return nil, status.Error(codes.AlreadyExists, "用户已存在")
	}

	user = model.User{Nickname: info.Nickname, Password: library.MD5(info.Password), Mobile: info.Mobile}
	result := db.Create(&user)
	if result.Error != nil {
		return nil, status.Error(codes.Unknown, "创建用户失败,未知原因")
	}
	userInfoResp := convertModelUserToResponseUser(user)
	return userInfoResp, nil

}

func (u *UserService) UpdateUser(ctx context.Context, info *proto.UpdateUserInfoRequest) (*emptypb.Empty, error) {
	var user model.User
	db, err := database.GetDB()
	if err != nil {
		zap.S().Errorf("数据库错误:%s\n", err.Error())
		return nil, status.Error(codes.DataLoss, "获取数据库出错")
	}
	db.First(&user, info.Id)
	if user.ID < 1 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	user.Nickname = info.Nickname
	user.Gender = uint8(info.Gender)
	user.Birthday = int(info.Birthday)

	result := db.Save(&user)
	if result.Error != nil {
		return nil, status.Error(codes.Unknown, "更新用户失败,未知原因")
	}
	return &emptypb.Empty{}, nil
}

func (u *UserService) GetUserList(ctx context.Context, request *proto.PageInfoRequest) (resp *proto.UserListResponse, err error) {
	var users []model.User
	db, err := database.GetDB()
	if err != nil {
		zap.S().Errorf("数据库错误:%s\n", err.Error())
		return nil, status.Error(codes.DataLoss, "获取数据库出错")
	}
	var count int64
	db.Model(&users).Count(&count)
	resp = &proto.UserListResponse{}
	resp.Total = int32(count)

	var page uint32 = 1
	var pageNum uint32 = 10
	if request.PageSize > 0 {
		pageNum = request.PageSize
	}
	if request.PageNum > 1 {
		page = request.PageNum
	}
	offset := pageNum * (page - 1)
	db.Offset(int(offset)).Limit(int(pageNum)).Find(&users)
	for _, value := range users {
		userInfoResp := convertModelUserToResponseUser(value)
		resp.Data = append(resp.Data, userInfoResp)
	}
	return resp, nil
}

func (u *UserService) CheckPassword(ctx context.Context, request *proto.CheckPasswordRequest) (resp *proto.CheckPasswordResponse, err error) {
	md5Pwd := library.MD5(request.Password)
	resp = &proto.CheckPasswordResponse{}
	resp.Success = false
	if request.Md5Password == md5Pwd {
		resp.Success = true
		return resp, nil
	}
	return resp, nil
}
