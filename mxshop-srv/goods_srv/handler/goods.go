package handler

import (
	"context"
	"go.uber.org/zap"
	"goods_srv/database"
	"goods_srv/model"
	"goods_srv/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GoodsService struct{}

func (g GoodsService) GoodsList(ctx context.Context, request *proto.GoodsFilterRequest) (resp *proto.GoodsListResponse, err error) {
	var goods []model.Goods
	db, err := database.GetDB()
	if err != nil {
		zap.S().Errorf("获取数据库失败:%s", err.Error())
		return nil, status.Error(codes.DataLoss, "获取数据库出错")
	}
	//var goodsList []model.Goods
	buildDB := db.Model(&model.Goods{})
	if request.KeyWords != "" {
		buildDB = buildDB.Where("name LIKE ?", "%"+request.KeyWords+"%")
	}
	if request.IsHot {
		buildDB = buildDB.Where("is_hot = 1")
	}
	if request.IsNew {
		buildDB = buildDB.Where("is_new = 1")
	}
	if request.PriceMin > 0 {
		buildDB = buildDB.Where("shop_price >= ?", request.PriceMin)
	}
	if request.PriceMax > 0 {
		buildDB = buildDB.Where("shop_price >= ?", request.PriceMax)
	}
	if request.Brand > 0 {
		buildDB = buildDB.Where("brand = ?", request.Brand)
	}
	if request.TopCategory > 0 { //未实现具体逻辑,无数据库
		var category model.Category
		db.Model(&model.Category{}).First(&category, request.TopCategory)
		if category.Level == 1 {

		} else if category.Level == 2 {

		} else if category.Level == 3 {

		}
	}

	//var page int32 = 1
	//var pageNum int32 = 10
	//if request.PagePerNums > 0 {
	//	pageNum = request.PagePerNums
	//}
	//if request.Pages > 1 {
	//	page = request.Pages
	//}
	var count int64
	db.Model(&goods).Count(&count)
	resp = &proto.GoodsListResponse{}

	resp.Total = int32(count)
	//for _, value :=range users {
	//	userInfoResp := convertModelUserToResponseUser(value)
	//	resp.Data = append(resp.Data,userInfoResp)
	//}
	return resp, nil
}

func (g GoodsService) BatchGetGoods(ctx context.Context, info *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	panic("implement me")
}

func (g GoodsService) CreateGoods(ctx context.Context, info *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	panic("implement me")
}

func (g GoodsService) DeleteGoods(ctx context.Context, info *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g GoodsService) UpdateGoods(ctx context.Context, info *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g GoodsService) GetGoodsDetail(ctx context.Context, request *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	panic("implement me")
}

func (g GoodsService) GetAllCategorysList(ctx context.Context, empty *emptypb.Empty) (*proto.CategoryListResponse, error) {
	panic("implement me")
}

func (g GoodsService) GetSubCategory(ctx context.Context, request *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	panic("implement me")
}

func (g GoodsService) CreateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	panic("implement me")
}

func (g GoodsService) DeleteCategory(ctx context.Context, request *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g GoodsService) UpdateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g GoodsService) BrandList(ctx context.Context, request *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	panic("implement me")
}

func (g GoodsService) CreateBrand(ctx context.Context, request *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	panic("implement me")
}

func (g GoodsService) DeleteBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g GoodsService) UpdateBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g GoodsService) BannerList(ctx context.Context, empty *emptypb.Empty) (*proto.BannerListResponse, error) {
	panic("implement me")
}

func (g GoodsService) CreateBanner(ctx context.Context, request *proto.BannerRequest) (*proto.BannerResponse, error) {
	panic("implement me")
}

func (g GoodsService) DeleteBanner(ctx context.Context, request *proto.BannerRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g GoodsService) UpdateBanner(ctx context.Context, request *proto.BannerRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g GoodsService) CategoryBrandList(ctx context.Context, request *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	panic("implement me")
}

func (g GoodsService) GetCategoryBrandList(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	panic("implement me")
}

func (g GoodsService) CreateCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	panic("implement me")
}

func (g GoodsService) DeleteCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g GoodsService) UpdateCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	panic("implement me")
}
