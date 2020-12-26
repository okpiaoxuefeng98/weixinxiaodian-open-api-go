# SkuUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int64** | 小商店内部商品ID，与out_product_id二选一 | [optional] 
**OutProductId** | **string** | 商家自定义商品ID，与product_id二选一 | [optional] 
**SkuId** | **int64** | out_sku_id、sku_id 二选一 | [optional] 
**OutSkuId** | **int64** | out_sku_id、sku_id二选一 | [optional] 
**ThumbImg** | **string** | sku小图 | 
**SalePrice** | **int64** | 售卖价格,以分为单位 | 
**MarketPrice** | **int64** | 市场价格,以分为单位 | [optional] 
**StockNum** | **int64** | 库存 | 
**SkuCode** | **string** | 商品编码 | [optional] 
**Barcode** | **string** | 条形码 | [optional] 
**SkuAttrs** | [**[]Attrs**](Attrs.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


