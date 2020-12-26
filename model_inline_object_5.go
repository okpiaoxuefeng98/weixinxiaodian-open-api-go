/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	// 小商店内部商品ID，与out_product_id二选一
	ProductId int64 `json:"product_id,omitempty"`
	// 商家自定义商品ID，与product_id二选一
	OutProductId string `json:"out_product_id,omitempty"`
}