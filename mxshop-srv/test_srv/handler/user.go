package handler

import (
	"context"
	"test_srv/proto"
)

type UserService struct{}

func (u UserService) GetUserList(ctx context.Context, request *proto.PageInfoRequest) (resp *proto.UserListResponse, err error) {
	resp = &proto.UserListResponse{}
	return resp, err
}
