# \RegisterApi

All URIs are relative to *https://api.weixin.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAuditStatus**](RegisterApi.md#CheckAuditStatus) | **Post** /product/register/check_audit_status | 异步状态查询
[**RegisterShop**](RegisterApi.md#RegisterShop) | **Post** /product/register/register_shop | 注册小商店账号
[**SubmitBasicinfo**](RegisterApi.md#SubmitBasicinfo) | **Post** /product/register/submit_basicinfo | 提交小商店基础信息
[**SubmitMerchantinfo**](RegisterApi.md#SubmitMerchantinfo) | **Post** /product/register/submit_merchantinfo | 提交支付资质



## CheckAuditStatus

> CheckAuditStatusRsp CheckAuditStatus(ctx, accessToken, body)

异步状态查询

异步状态查询

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**CheckAuditStatusRsp**](CheckAuditStatusRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterShop

> RegisterShopRsp RegisterShop(ctx, accessToken, body)

注册小商店账号

注册小商店账号

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject1**](InlineObject1.md)|  | 

### Return type

[**RegisterShopRsp**](RegisterShopRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBasicinfo

> SubmitBasicinfo SubmitBasicinfo(ctx, accessToken, body)

提交小商店基础信息

提交小商店基础信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject3**](InlineObject3.md)|  | 

### Return type

[**SubmitBasicinfo**](SubmitBasicinfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitMerchantinfo

> SubmitMerchantinfo SubmitMerchantinfo(ctx, accessToken, body)

提交支付资质

提交支付资质

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject2**](InlineObject2.md)|  | 

### Return type

[**SubmitMerchantinfo**](SubmitMerchantinfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

