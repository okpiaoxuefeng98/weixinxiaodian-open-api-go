/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// CheckAuditStatusRspDataPayAuditDetail struct for CheckAuditStatusRspDataPayAuditDetail
type CheckAuditStatusRspDataPayAuditDetail struct {
	// 参数名
	ParamName string `json:"param_name,omitempty"`
	// 支付资质驳回原因
	RejectReason string `json:"reject_reason,omitempty"`
}
