/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// DeliveryInfo 
type DeliveryInfo struct {
	// 快递方式（目前只有\"快递\"）
	DeliveryMethod string `json:"delivery_method,omitempty"`
	// 发货时间
	DeliveryTime string `json:"delivery_time,omitempty"`
	DeliveryProductInfo DeliveryInfoDeliveryProductInfo `json:"delivery_product_info,omitempty"`
	AddressInfo AddressInfo `json:"address_info,omitempty"`
}