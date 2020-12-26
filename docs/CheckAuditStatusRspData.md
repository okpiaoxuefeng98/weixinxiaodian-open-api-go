# CheckAuditStatusRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisterStatus** | **int32** | 注册状态 0-成功 1-已发送协议还未签约 2-未发送协议或协议已过期，需发送协议，当register_status为0时以下字段有意义 | [optional] 
**MerchantInfoStatus** | **int32** | 商家信息状态, 具体含义查看状态枚举值 | [optional] 
**AcctVerifyStatus** | **int32** | 账户验证状态, 具体含义查看状态枚举值 | [optional] 
**BasicInfoStatus** | **int32** | 基础信息状态, 具体含义查看状态枚举值 | [optional] 
**PaySignStatus** | **int32** | 支付签约状态, 具体含义查看状态枚举值 | [optional] 
**AuditRejectReasons** | **string** | 基础信息驳回原因 | [optional] 
**LegalValidationUrl** | **string** | 法人验证链接 | [optional] 
**PayAuditDetail** | [**CheckAuditStatusRspDataPayAuditDetail**](CheckAuditStatusRsp_data_pay_audit_detail.md) |  | [optional] 
**RegisteredAppid** | **string** | 注册的appid | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


