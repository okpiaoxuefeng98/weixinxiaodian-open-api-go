/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// SubmitBasicinfo 
type SubmitBasicinfo struct {
	// 系统异常--1;token太长--2;-15-缺少类目资质
	Errcode int32 `json:"errcode,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
}