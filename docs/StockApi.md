# \StockApi

All URIs are relative to *https://api.weixin.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StockGet**](StockApi.md#StockGet) | **Post** /product/stock/get | 获取库存
[**StockUpdate**](StockApi.md#StockUpdate) | **Post** /product/stock/update | 更新库存



## StockGet

> StockGetRsp StockGet(ctx, accessToken, body)

获取库存

获取库存

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject16**](InlineObject16.md)|  | 

### Return type

[**StockGetRsp**](StockGetRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StockUpdate

> StockUpdateRsp StockUpdate(ctx, accessToken, body)

更新库存

更新库存

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject15**](InlineObject15.md)|  | 

### Return type

[**StockUpdateRsp**](StockUpdateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

