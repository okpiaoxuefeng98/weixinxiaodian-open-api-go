/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// OrderGetListRsp 
type OrderGetListRsp struct {
	// 系统异常--1;product_id不存在-100002;sku不存在-100004
	Errcode int32 `json:"errcode,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	TotalNum int64 `json:"total_num,omitempty"`
	Orders []Order `json:"orders,omitempty"`
}
