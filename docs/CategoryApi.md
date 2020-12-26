# \CategoryApi

All URIs are relative to *https://api.weixin.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoryGet**](CategoryApi.md#CategoryGet) | **Post** /product/category/get | 获取类目详情



## CategoryGet

> CategoryGetRsp CategoryGet(ctx, accessToken, body)

获取类目详情

获取类目详情

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject4**](InlineObject4.md)|  | 

### Return type

[**CategoryGetRsp**](CategoryGetRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

