/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// InlineObject20 struct for InlineObject20
type InlineObject20 struct {
	// 订单id
	OrderId string `json:"order_id"`
	DeliveryList []ProductDeliverySendDeliveryList `json:"delivery_list"`
}
