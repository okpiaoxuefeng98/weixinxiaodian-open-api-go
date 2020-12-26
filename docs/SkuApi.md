# \SkuApi

All URIs are relative to *https://api.weixin.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SkuAdd**](SkuApi.md#SkuAdd) | **Post** /product/sku/add | 添加SKU
[**SkuBatchAdd**](SkuApi.md#SkuBatchAdd) | **Post** /product/sku/batch_add | 批量添加SKU
[**SkuDel**](SkuApi.md#SkuDel) | **Post** /product/sku/del | 删除SKU
[**SkuGet**](SkuApi.md#SkuGet) | **Post** /product/sku/get | 获取商品
[**SkuGetList**](SkuApi.md#SkuGetList) | **Post** /product/sku/get_list | 批量获取SKU信息
[**SkuUpdate**](SkuApi.md#SkuUpdate) | **Post** /product/sku/update | 更新商品
[**SkuUpdatePrice**](SkuApi.md#SkuUpdatePrice) | **Post** /product/sku/update_price | 更新SKU价格



## SkuAdd

> SkuAddRsp SkuAdd(ctx, accessToken, body)

添加SKU

添加SKU

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | **Sku**|  | 

### Return type

[**SkuAddRsp**](SkuAddRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkuBatchAdd

> SkuBatchAddRsp SkuBatchAdd(ctx, accessToken, body)

批量添加SKU

批量添加SKU

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**SkuBatchAddReq**](SkuBatchAddReq.md)|  | 

### Return type

[**SkuBatchAddRsp**](SkuBatchAddRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkuDel

> SkuDelRsp SkuDel(ctx, accessToken, body)

删除SKU

删除SKU

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject11**](InlineObject11.md)|  | 

### Return type

[**SkuDelRsp**](SkuDelRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkuGet

> SkuGetRsp SkuGet(ctx, accessToken, body)

获取商品

获取商品

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject12**](InlineObject12.md)|  | 

### Return type

[**SkuGetRsp**](SkuGetRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkuGetList

> SkuGetListRsp SkuGetList(ctx, accessToken, body)

批量获取SKU信息

批量获取SKU信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject13**](InlineObject13.md)|  | 

### Return type

[**SkuGetListRsp**](SkuGetListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkuUpdate

> SkuUpdateRsp SkuUpdate(ctx, accessToken, body)

更新商品

更新商品

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**SkuUpdateReq**](SkuUpdateReq.md)|  | 

### Return type

[**SkuUpdateRsp**](SkuUpdateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkuUpdatePrice

> SkuUpdatePriceRsp SkuUpdatePrice(ctx, accessToken, body)

更新SKU价格

更新SKU价格

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject14**](InlineObject14.md)|  | 

### Return type

[**SkuUpdatePriceRsp**](SkuUpdatePriceRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

