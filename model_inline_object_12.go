/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// InlineObject12 struct for InlineObject12
type InlineObject12 struct {
	// out_sku_id、sku_id 二选一
	SkuId int64 `json:"sku_id,omitempty"`
	// out_sku_id、sku_id 二选一
	OutSkuId string `json:"out_sku_id,omitempty"`
	// 默认0:获取线上数据, 1:获取草稿数据
	NeedEditSku int32 `json:"need_edit_sku,omitempty"`
	// 默认0:获取草稿库存, 1:获取线上真实库存
	NeedRealStock int32 `json:"need_real_stock,omitempty"`
}
