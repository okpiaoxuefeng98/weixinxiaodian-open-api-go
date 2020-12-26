# \OrderApi

All URIs are relative to *https://api.weixin.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderGet**](OrderApi.md#OrderGet) | **Post** /product/order/get | 获取订单详情
[**OrderGetList**](OrderApi.md#OrderGetList) | **Post** /product/order/get_list | 获取订单列表
[**OrderSearch**](OrderApi.md#OrderSearch) | **Post** /product/order/search | 搜索订单



## OrderGet

> OrderGetRsp OrderGet(ctx, accessToken, body)

获取订单详情

获取订单详情

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject18**](InlineObject18.md)|  | 

### Return type

[**OrderGetRsp**](OrderGetRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetList

> OrderGetListRsp OrderGetList(ctx, accessToken, body)

获取订单列表

获取订单列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject17**](InlineObject17.md)|  | 

### Return type

[**OrderGetListRsp**](OrderGetListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderSearch

> OrderSearchRsp OrderSearch(ctx, accessToken, body)

搜索订单

搜索订单

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject19**](InlineObject19.md)|  | 

### Return type

[**OrderSearchRsp**](OrderSearchRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

