package model

type Goods struct {
	ID              int     `json:"id" gorm:"column:id"`
	Category        int     `json:"category" gorm:"column:category"`                   // 商品类目
	GoodsSn         string  `json:"goods_sn" gorm:"column:goods_sn"`                   // 商品唯一货号
	Name            string  `json:"name" gorm:"column:name"`                           // 商品名
	ClickNum        int     `json:"click_num" gorm:"column:click_num"`                 // 点击数
	SoldNum         int     `json:"sold_num" gorm:"column:sold_num"`                   // 商品销售量
	FavNum          int     `json:"fav_num" gorm:"column:fav_num"`                     // 收藏数
	MarketPrice     float64 `json:"market_price" gorm:"column:market_price"`           // 市场价格
	ShopPrice       float64 `json:"shop_price" gorm:"column:shop_price"`               // 本店价格
	GoodsBrief      string  `json:"goods_brief" gorm:"column:goods_brief"`             // 商品简短描述
	ShipFree        int8    `json:"ship_free" gorm:"column:ship_free"`                 // 是否承担运费
	Brand           int     `json:"brand" gorm:"column:brand"`                         // 品牌
	OnSale          uint    `json:"on_sale" gorm:"column:on_sale"`                     // 是否上架
	Images          string  `json:"images" gorm:"column:images"`                       // 商品轮播图
	DescImages      string  `json:"desc_images" gorm:"column:desc_images"`             // 详情页图片
	IsNew           int8    `json:"is_new" gorm:"column:is_new"`                       // 是否新品
	IsHot           int8    `json:"is_hot" gorm:"column:is_hot"`                       // 是否热销
	AddTime         int     `json:"add_time" gorm:"column:add_time"`                   // 添加时间
	UpdateTime      int     `json:"update_time" gorm:"column:update_time"`             // 更新时间
	IsDeleted       int8    `json:"is_deleted" gorm:"column:is_deleted"`               // 是否删除
	GoodsFrontImage string  `json:"goods_front_image" gorm:"column:goods_front_image"` // 封面图
}

func (m *Goods) TableName() string {
	return "goods"
}
