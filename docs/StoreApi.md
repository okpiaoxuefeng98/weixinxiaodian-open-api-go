# \StoreApi

All URIs are relative to *https://api.weixin.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StoreGetInfo**](StoreApi.md#StoreGetInfo) | **Post** /product/store/get_info | 获取店铺基本信息
[**StoreGetShopcat**](StoreApi.md#StoreGetShopcat) | **Post** /product/store/get_shopcat | 获取店铺的商品分类



## StoreGetInfo

> StoreGetInfoRsp StoreGetInfo(ctx, accessToken)

获取店铺基本信息

获取店铺基本信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 

### Return type

[**StoreGetInfoRsp**](StoreGetInfoRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreGetShopcat

> StoreGetShopcatRsp StoreGetShopcat(ctx, accessToken)

获取店铺的商品分类

获取店铺的商品分类

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 

### Return type

[**StoreGetShopcatRsp**](StoreGetShopcatRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

