/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// CouponInfoDiscountInfoDiscountCondition struct for CouponInfoDiscountInfoDiscountCondition
type CouponInfoDiscountInfoDiscountCondition struct {
	// 商品折扣券打折金额
	ProductCnt int64 `json:"product_cnt,omitempty"`
	// 商品id，商品折扣券需填写
	ProductIds int64 `json:"product_ids,omitempty"`
	// 商品价格，满减券需填写
	ProductPrice int64 `json:"product_price,omitempty"`
}