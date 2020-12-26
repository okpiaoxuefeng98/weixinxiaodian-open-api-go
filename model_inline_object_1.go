/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	// 微信号
	WxName string `json:"wx_name"`
	// 身份证姓名
	IdCardName string `json:"id_card_name"`
	// 身份证号
	IdCardNumber string `json:"id_card_number"`
	// 1-整店打包，2-组件开放
	ApiOpenstoreType int32 `json:"api_openstore_type"`
	// 渠道号
	ChannelId int32 `json:"channel_id,omitempty"`
}