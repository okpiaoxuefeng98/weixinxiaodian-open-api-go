/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// CouponPushRsp 
type CouponPushRsp struct {
	// 系统异常--1;token太长--2;参数有误-9401020;优惠券状态不对-109100;优惠券库存不足-109101;优惠券还没过期-109102;优惠券限领-109103
	Errcode int32 `json:"errcode,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
}
