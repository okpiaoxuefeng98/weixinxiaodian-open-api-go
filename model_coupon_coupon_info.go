/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// CouponCouponInfo 主图,多张,列表
type CouponCouponInfo struct {
	// 优惠券名称
	Name string `json:"name,omitempty"`
	ValidInfo CouponCouponInfoValidInfo `json:"valid_info,omitempty"`
}
