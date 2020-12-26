/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// Sku Sku
type Sku struct {
	// 小商店内部商品ID
	ProductId int64 `json:"product_id"`
	// 商家自定义商品ID
	OutProductId string `json:"out_product_id"`
	// 商家自定义skuID
	OutSkuId string `json:"out_sku_id"`
	// sku_id
	SkuId int64 `json:"sku_id"`
	// sku小图
	ThumbImg string `json:"thumb_img"`
	// 售卖价格,以分为单位
	SalePrice int64 `json:"sale_price"`
	// 市场价格,以分为单位
	MarketPrice int64 `json:"market_price"`
	// 库存
	StockNum int64 `json:"stock_num"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty"`
	// 条形码
	Barcode string `json:"barcode,omitempty"`
	// sku状态
	Status int32 `json:"status,omitempty"`
	// sku属性值
	SkuAttrs []Attrs `json:"sku_attrs"`
}