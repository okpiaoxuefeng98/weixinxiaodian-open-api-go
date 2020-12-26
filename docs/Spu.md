# Spu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | 标题 | [optional] 
**SubTitle** | **string** | 副标题 | [optional] 
**HeadImg** | **[]string** | 主图,多张,列表 | [optional] 
**DescInfo** | [**SpuDescInfo**](Spu_desc_info.md) |  | [optional] 
**OutProductId** | **string** | 商家自定义商品ID | [optional] 
**ProductId** | **int64** | 小商店内部商品ID | [optional] 
**BrandId** | **int64** | 商家需要申请品牌 | [optional] 
**Status** | **int64** | 商品线上状态;初始值-0;上架-5;回收站-6;逻辑删除-9;自主下架-11;售磬下架-12;违规下架/风控系统下架-13 | [optional] 
**EditStatus** | **int64** | 商品草稿状态;初始值-0;编辑中-1;审核中-2;审核失败-3;审核成功-4 | [optional] 
**MinPrice** | **int64** | 商品SKU最小价格（单位：分） | [optional] 
**Cats** | [**[]Cats**](Cats.md) | 类目 | [optional] 
**Attrs** | [**[]Attrs**](Attrs.md) | 属性 | [optional] 
**Model** | **string** | 商品型号 | [optional] 
**Shopcat** | [**[]SpuShopcat**](Spu_shopcat.md) | 该 skus 列表非必填，可另行通过 BatchAddSKU 添加 SKU | [optional] 
**Skus** | [**[]SpuSkus**](Spu_skus.md) | 该 skus 列表非必填，可另行通过 BatchAddSKU 添加 SKU | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


