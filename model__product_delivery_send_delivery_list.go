/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// ProductDeliverySendDeliveryList struct for ProductDeliverySendDeliveryList
type ProductDeliverySendDeliveryList struct {
	// 快递公司id
	DeliveryId string `json:"delivery_id"`
	// 快递单号
	WaybillId string `json:"waybill_id"`
}