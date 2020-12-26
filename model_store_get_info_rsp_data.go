/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// StoreGetInfoRspData struct for StoreGetInfoRspData
type StoreGetInfoRspData struct {
	// 商店名称
	StoreName string `json:"store_name,omitempty"`
	// 商店头像
	Logo string `json:"logo,omitempty"`
	// 1-企业店,;2-个人店
	Type int32 `json:"type,omitempty"`
}
