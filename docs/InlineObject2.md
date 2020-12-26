# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appid** | **string** |  | 
**SubjectType** | **string** | 主体类型,\&quot;4\&quot;-个体工商户，营业执照上的主体类型一般为个体户、个体工商户、个体经营。\&quot;2\&quot;-企业，营业执照上的主体类型一般为有限公司、有限责任公司。 | 
**BusiLicense** | [**BusiLicense**](BusiLicense.md) |  | 
**OrganizationCodeInfo** | [**OrganizationCodeInfo**](OrganizationCodeInfo.md) |  | [optional] 
**IdCardInfo** | [**IdCardInfo**](IdCardInfo.md) |  | 
**SuperAdministratorInfo** | [**SuperAdministratorInfo**](SuperAdministratorInfo.md) |  | 
**MerchantShortname** | **string** | 商户简称 UTF-8格式，中文占3个字节，即最多16个汉字长度。将在支付完成页向买家展示，需与商家的实际售卖商品相符 。示例值：腾讯 | 
**SpecialQualification** | [**[]SuperAdministratorInfo**](SuperAdministratorInfo.md) | 特殊资质 1、根据商户经营业务要求提供相关资质，详情查看《行业对应特殊资质》。2、请提供为“申请商家主体”所属的特殊资质，可授权使用总公司/分公司的特殊资 质；3、最多可上传5张照片，请填写通过图片上传接口预先上传图片生成好的MediaID 。 | [optional] 
**SupplementaryMaterial** | [**[]SuperAdministratorInfo**](SuperAdministratorInfo.md) | 补充材料 | [optional] 
**SupplementaryDesc** | **string** | 补充描述 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


