/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// InlineObject7 struct for InlineObject7
type InlineObject7 struct {
	// 商品状态
	Status int32 `json:"status"`
	// 第几页（最小填1）
	Page int32 `json:"page"`
	// 每页数量(不超过10,000)
	PageSize int32 `json:"page_size"`
	// 默认0:获取线上数据, 1:获取草稿数据
	NeedEditSpu int32 `json:"need_edit_spu,omitempty"`
}