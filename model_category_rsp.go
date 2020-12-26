/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// CategoryRsp 类目详情
type CategoryRsp struct {
	// 类目ID
	CatId int64 `json:"cat_id,omitempty"`
	// 类目父ID
	FCatId int64 `json:"f_cat_id,omitempty"`
	// 类目名称
	Name string `json:"name,omitempty"`
}