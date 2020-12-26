# \SpuApi

All URIs are relative to *https://api.weixin.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpuAdd**](SpuApi.md#SpuAdd) | **Post** /product/spu/add | 添加商品
[**SpuDel**](SpuApi.md#SpuDel) | **Post** /product/spu/del | 删除商品
[**SpuDelisting**](SpuApi.md#SpuDelisting) | **Post** /product/spu/delisting | 下架商品
[**SpuGet**](SpuApi.md#SpuGet) | **Post** /product/spu/get | 获取商品
[**SpuGetList**](SpuApi.md#SpuGetList) | **Post** /product/spu/get_list | 获取商品列表
[**SpuListing**](SpuApi.md#SpuListing) | **Post** /product/spu/listing | 上架商品
[**SpuSearch**](SpuApi.md#SpuSearch) | **Post** /product/spu/search | 搜索商品
[**SpuUpdate**](SpuApi.md#SpuUpdate) | **Post** /product/spu/update | 更新商品



## SpuAdd

> SpuAddRsp SpuAdd(ctx, accessToken, body)

添加商品

添加商品

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**SpuAddReq**](SpuAddReq.md)|  | 

### Return type

[**SpuAddRsp**](SpuAddRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpuDel

> SpuDelRsp SpuDel(ctx, accessToken, body)

删除商品

删除商品

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject5**](InlineObject5.md)|  | 

### Return type

[**SpuDelRsp**](SpuDelRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpuDelisting

> SpuDelistingRsp SpuDelisting(ctx, accessToken, body)

下架商品

下架商品

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject10**](InlineObject10.md)|  | 

### Return type

[**SpuDelistingRsp**](SpuDelistingRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpuGet

> SpuGetRsp SpuGet(ctx, accessToken, body)

获取商品

获取商品

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject6**](InlineObject6.md)|  | 

### Return type

[**SpuGetRsp**](SpuGetRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpuGetList

> SpuGetListRsp SpuGetList(ctx, accessToken, body)

获取商品列表

获取商品列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject7**](InlineObject7.md)|  | 

### Return type

[**SpuGetListRsp**](SpuGetListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpuListing

> SpuListingRsp SpuListing(ctx, accessToken, body)

上架商品

上架商品

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject9**](InlineObject9.md)|  | 

### Return type

[**SpuListingRsp**](SpuListingRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpuSearch

> SpuSearchRsp SpuSearch(ctx, accessToken, body)

搜索商品

搜索商品

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject8**](InlineObject8.md)|  | 

### Return type

[**SpuSearchRsp**](SpuSearchRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpuUpdate

> SpuUpdateRsp SpuUpdate(ctx, accessToken, body)

更新商品

更新商品

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**SpuUpdateReq**](SpuUpdateReq.md)|  | 

### Return type

[**SpuUpdateRsp**](SpuUpdateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

