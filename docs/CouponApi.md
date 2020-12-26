# \CouponApi

All URIs are relative to *https://api.weixin.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CouponCreate**](CouponApi.md#CouponCreate) | **Post** /product/coupon/create | 创建优惠券
[**CouponGetList**](CouponApi.md#CouponGetList) | **Post** /product/coupon/get_list | 获取优惠券列表
[**CouponPush**](CouponApi.md#CouponPush) | **Post** /product/coupon/push | 发放优惠券



## CouponCreate

> CouponGetListRsp CouponCreate(ctx, accessToken, body)

创建优惠券

创建优惠券

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject21**](InlineObject21.md)|  | 

### Return type

[**CouponGetListRsp**](CouponGetListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CouponGetList

> CouponGetListRsp CouponGetList(ctx, accessToken, body)

获取优惠券列表

获取优惠券列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject22**](InlineObject22.md)|  | 

### Return type

[**CouponGetListRsp**](CouponGetListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CouponPush

> CouponPushRsp CouponPush(ctx, accessToken, body)

发放优惠券

发放优惠券

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject23**](InlineObject23.md)|  | 

### Return type

[**CouponPushRsp**](CouponPushRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

