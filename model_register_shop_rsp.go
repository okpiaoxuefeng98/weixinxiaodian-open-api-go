/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// RegisterShopRsp 
type RegisterShopRsp struct {
	// 系统异常--1;token太长--2;1005-实名认证失败，请检查身份证、微信号和姓名是否匹配
	Errcode int32 `json:"errcode,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
}
