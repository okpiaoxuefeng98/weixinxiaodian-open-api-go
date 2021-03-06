/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// Order 
type Order struct {
	// 订单ID
	OrderId int64 `json:"order_id,omitempty"`
	// 订单状态
	Status int32 `json:"status,omitempty"`
	// 创建时间
	CraeteTime string `json:"craete_time,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty"`
	OrderDetail OrderOrderDetail `json:"order_detail,omitempty"`
	AfterSaleInfo AfterSaleInfo `json:"after_sale_info,omitempty"`
	// 用户的openid，用于物流助手接口
	Openid string `json:"openid,omitempty"`
	ExtInfo OrderExtInfo `json:"ext_info,omitempty"`
}
