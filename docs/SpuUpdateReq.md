# SpuUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int64** | 小商店内部商品ID，与out_product_id二选一 | [optional] 
**OutProductId** | **string** | 商家自定义商品ID，与product_id二选一 | [optional] 
**Title** | **string** | 标题 | 
**SubTitle** | **string** | 副标题 | [optional] 
**HeadImg** | **[]string** | 主图,多张,列表 | 
**DescInfo** | [**SpuAddReqDescInfo**](SpuAddReq_desc_info.md) |  | 
**BrandId** | **int64** | 品牌ID，商家需要申请品牌并通过获取品牌接口brand/get获取，如果是无品牌请填2100000000 | 
**Cats** | [**[]Cats**](Cats.md) | 类目 | 
**Attrs** | [**[]Attrs**](Attrs.md) | 属性 | 
**Model** | **string** | 商品型号 | [optional] 
**ExpressInfo** | [**SpuUpdateReqExpressInfo**](SpuUpdateReq_express_info.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


