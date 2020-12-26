/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// CouponInfo 
type CouponInfo struct {
	// 优惠券名称
	Name string `json:"name"`
	DiscountInfo CouponInfoDiscountInfo `json:"discount_info"`
	ExtInfo CouponInfoExtInfo `json:"ext_info"`
	PromoteInfo CouponInfoPromoteInfo `json:"promote_info"`
	ReceiveInfo CouponInfoReceiveInfo `json:"receive_info"`
	ValidInfo CouponInfoValidInfo `json:"valid_info"`
}
