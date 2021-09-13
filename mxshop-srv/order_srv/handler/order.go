package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"order_srv/proto"
)

type OrderService struct{}

func (o OrderService) CartItemList(ctx context.Context, info *proto.UserInfo) (*proto.CartItemListResponse, error) {
	panic("implement me")
}

func (o OrderService) CreateCartItem(ctx context.Context, request *proto.CartItemRequest) (*proto.ShopCartInfoResponse, error) {
	panic("implement me")
}

func (o OrderService) UpdateCartItem(ctx context.Context, request *proto.CartItemRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (o OrderService) DeleteCartItem(ctx context.Context, request *proto.CartItemRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (o OrderService) CreateOrder(ctx context.Context, request *proto.OrderRequest) (*proto.OrderInfoResponse, error) {
	panic("implement me")
}

func (o OrderService) OrderList(ctx context.Context, request *proto.OrderFilterRequest) (*proto.OrderListResponse, error) {
	panic("implement me")
}

func (o OrderService) OrderDetail(ctx context.Context, request *proto.OrderRequest) (*proto.OrderInfoDetailResponse, error) {
	panic("implement me")
}

func (o OrderService) UpdateOrderStatus(ctx context.Context, status *proto.OrderStatus) (*emptypb.Empty, error) {
	panic("implement me")
}
