/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// SpuAddReqDescInfo 商品详情，图文(目前只支持图片)
type SpuAddReqDescInfo struct {
	Imgs []string `json:"imgs"`
}