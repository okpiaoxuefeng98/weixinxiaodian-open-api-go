/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// SubmitMerchantinfo 
type SubmitMerchantinfo struct {
	// 系统异常--1;token太长--2;40-参数错误,请使用正确的参数重新调用;429-请降低调用频率;1005-实名认证失败，请检查身份证、微信号和姓名是否匹配
	Errcode int32 `json:"errcode,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
}