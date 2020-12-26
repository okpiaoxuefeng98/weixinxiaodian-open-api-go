/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// SpuUpdateReq 
type SpuUpdateReq struct {
	// 小商店内部商品ID，与out_product_id二选一
	ProductId int64 `json:"product_id,omitempty"`
	// 商家自定义商品ID，与product_id二选一
	OutProductId string `json:"out_product_id,omitempty"`
	// 标题
	Title string `json:"title"`
	// 副标题
	SubTitle string `json:"sub_title,omitempty"`
	// 主图,多张,列表
	HeadImg []string `json:"head_img"`
	DescInfo SpuAddReqDescInfo `json:"desc_info"`
	// 品牌ID，商家需要申请品牌并通过获取品牌接口brand/get获取，如果是无品牌请填2100000000
	BrandId int64 `json:"brand_id"`
	// 类目
	Cats []Cats `json:"cats"`
	// 属性
	Attrs []Attrs `json:"attrs"`
	// 商品型号
	Model string `json:"model,omitempty"`
	ExpressInfo SpuUpdateReqExpressInfo `json:"express_info"`
}
