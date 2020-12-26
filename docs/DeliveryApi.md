# \DeliveryApi

All URIs are relative to *https://api.weixin.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliveryGetCompanyList**](DeliveryApi.md#DeliveryGetCompanyList) | **Post** /product/delivery/get_company_list | 获取快递公司列表
[**DeliveryGetFreightTemplate**](DeliveryApi.md#DeliveryGetFreightTemplate) | **Post** /product/delivery/get_freight_template | 获取运费模板
[**DeliverySend**](DeliveryApi.md#DeliverySend) | **Post** /product/delivery/send | 订单发货



## DeliveryGetCompanyList

> DeliveryGetCompanyListRsp DeliveryGetCompanyList(ctx, accessToken)

获取快递公司列表

获取快递公司列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 

### Return type

[**DeliveryGetCompanyListRsp**](DeliveryGetCompanyListRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGetFreightTemplate

> DeliveryGetFreightTemplateRsp DeliveryGetFreightTemplate(ctx, accessToken)

获取运费模板

获取运费模板

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 

### Return type

[**DeliveryGetFreightTemplateRsp**](DeliveryGetFreightTemplateRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliverySend

> DeliverySendRsp DeliverySend(ctx, accessToken, body)

订单发货

订单发货

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessToken** | **string**| 访问权限token | 
**body** | [**InlineObject20**](InlineObject20.md)|  | 

### Return type

[**DeliverySendRsp**](DeliverySendRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

