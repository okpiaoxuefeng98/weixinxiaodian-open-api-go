# InlineObject15

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int64** | 小商店内部商品ID，与out_product_id二选一 | [optional] 
**OutProductId** | **string** | 商家自定义商品ID，与product_id二选一 | [optional] 
**SkuId** | **int64** | out_sku_id、sku_id 二选一 | [optional] 
**OutSkuId** | **string** | out_sku_id、sku_id二选一 | [optional] 
**Type** | **int32** | 1:全量更新 2:增量更新 | 
**StockNum** | **int64** | 全量更新时，stock_num必须大于0；增量更新时正数增加库存，负数减库存 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


