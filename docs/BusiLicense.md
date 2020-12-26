# BusiLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseType** | **int64** | 营业执照类型 1-三证合一 2-普通营业执照 | 
**PicFile** | [**PicFile**](PicFile.md) |  | 
**RegistrationNum** | **string** | 请填写营业执照上的注册号/统一社会信用代码，须为15位数字或18位数字大写字母。示例值;123456789012345678 特殊规则;长度最小15个字节 | 
**MerchantName** | **string** | 1、请填写营业执照/登记证书的商家名称，2~110个字符，支持括号 2、个体工商户/党政、机关及事业单位，不能以“公司”结尾。 3、个体工商户，若营业执照上商户名称为空或为“无”，请填写\&quot;个体户+经营者姓名\&quot;，如“个体户张三” 。示例值-腾讯科技有限公司 | 
**LegalRepresentative** | **string** | 请填写证件的经营者/法定代表人姓名。示例值-张三 | 
**RegisteredAddrs** | **string** | 注册地址 | [optional] 
**StartDate** | **string** | 注册日期，格式2014-01-01 | 
**EndDate** | **string** | 结束有效期，格式;2014-01-01;1、若证件有效期为长期，请填写;长期。2、结束时间需大于开始时间。3、有效期必须大于60天，即结束时间距当前时间需超过60天。 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


