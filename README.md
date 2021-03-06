# Go API client for wxxiaodian

wei xin xiao shang dian open api

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./wxxiaodian"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.weixin.qq.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BrandApi* | [**BrandGet**](docs/BrandApi.md#brandget) | **Post** /product/brand/get | 获取品牌列表
*CategoryApi* | [**CategoryGet**](docs/CategoryApi.md#categoryget) | **Post** /product/category/get | 获取类目详情
*CouponApi* | [**CouponCreate**](docs/CouponApi.md#couponcreate) | **Post** /product/coupon/create | 创建优惠券
*CouponApi* | [**CouponGetList**](docs/CouponApi.md#coupongetlist) | **Post** /product/coupon/get_list | 获取优惠券列表
*CouponApi* | [**CouponPush**](docs/CouponApi.md#couponpush) | **Post** /product/coupon/push | 发放优惠券
*DeliveryApi* | [**DeliveryGetCompanyList**](docs/DeliveryApi.md#deliverygetcompanylist) | **Post** /product/delivery/get_company_list | 获取快递公司列表
*DeliveryApi* | [**DeliveryGetFreightTemplate**](docs/DeliveryApi.md#deliverygetfreighttemplate) | **Post** /product/delivery/get_freight_template | 获取运费模板
*DeliveryApi* | [**DeliverySend**](docs/DeliveryApi.md#deliverysend) | **Post** /product/delivery/send | 订单发货
*OrderApi* | [**OrderGet**](docs/OrderApi.md#orderget) | **Post** /product/order/get | 获取订单详情
*OrderApi* | [**OrderGetList**](docs/OrderApi.md#ordergetlist) | **Post** /product/order/get_list | 获取订单列表
*OrderApi* | [**OrderSearch**](docs/OrderApi.md#ordersearch) | **Post** /product/order/search | 搜索订单
*RegisterApi* | [**CheckAuditStatus**](docs/RegisterApi.md#checkauditstatus) | **Post** /product/register/check_audit_status | 异步状态查询
*RegisterApi* | [**RegisterShop**](docs/RegisterApi.md#registershop) | **Post** /product/register/register_shop | 注册小商店账号
*RegisterApi* | [**SubmitBasicinfo**](docs/RegisterApi.md#submitbasicinfo) | **Post** /product/register/submit_basicinfo | 提交小商店基础信息
*RegisterApi* | [**SubmitMerchantinfo**](docs/RegisterApi.md#submitmerchantinfo) | **Post** /product/register/submit_merchantinfo | 提交支付资质
*SkuApi* | [**SkuAdd**](docs/SkuApi.md#skuadd) | **Post** /product/sku/add | 添加SKU
*SkuApi* | [**SkuBatchAdd**](docs/SkuApi.md#skubatchadd) | **Post** /product/sku/batch_add | 批量添加SKU
*SkuApi* | [**SkuDel**](docs/SkuApi.md#skudel) | **Post** /product/sku/del | 删除SKU
*SkuApi* | [**SkuGet**](docs/SkuApi.md#skuget) | **Post** /product/sku/get | 获取商品
*SkuApi* | [**SkuGetList**](docs/SkuApi.md#skugetlist) | **Post** /product/sku/get_list | 批量获取SKU信息
*SkuApi* | [**SkuUpdate**](docs/SkuApi.md#skuupdate) | **Post** /product/sku/update | 更新商品
*SkuApi* | [**SkuUpdatePrice**](docs/SkuApi.md#skuupdateprice) | **Post** /product/sku/update_price | 更新SKU价格
*SpuApi* | [**SpuAdd**](docs/SpuApi.md#spuadd) | **Post** /product/spu/add | 添加商品
*SpuApi* | [**SpuDel**](docs/SpuApi.md#spudel) | **Post** /product/spu/del | 删除商品
*SpuApi* | [**SpuDelisting**](docs/SpuApi.md#spudelisting) | **Post** /product/spu/delisting | 下架商品
*SpuApi* | [**SpuGet**](docs/SpuApi.md#spuget) | **Post** /product/spu/get | 获取商品
*SpuApi* | [**SpuGetList**](docs/SpuApi.md#spugetlist) | **Post** /product/spu/get_list | 获取商品列表
*SpuApi* | [**SpuListing**](docs/SpuApi.md#spulisting) | **Post** /product/spu/listing | 上架商品
*SpuApi* | [**SpuSearch**](docs/SpuApi.md#spusearch) | **Post** /product/spu/search | 搜索商品
*SpuApi* | [**SpuUpdate**](docs/SpuApi.md#spuupdate) | **Post** /product/spu/update | 更新商品
*StockApi* | [**StockGet**](docs/StockApi.md#stockget) | **Post** /product/stock/get | 获取库存
*StockApi* | [**StockUpdate**](docs/StockApi.md#stockupdate) | **Post** /product/stock/update | 更新库存
*StoreApi* | [**StoreGetInfo**](docs/StoreApi.md#storegetinfo) | **Post** /product/store/get_info | 获取店铺基本信息
*StoreApi* | [**StoreGetShopcat**](docs/StoreApi.md#storegetshopcat) | **Post** /product/store/get_shopcat | 获取店铺的商品分类


## Documentation For Models

 - [AddressInfo](docs/AddressInfo.md)
 - [AfterSaleInfo](docs/AfterSaleInfo.md)
 - [AfterSaleInfoAftersaleOrderList](docs/AfterSaleInfoAftersaleOrderList.md)
 - [Attrs](docs/Attrs.md)
 - [BrandGetRsp](docs/BrandGetRsp.md)
 - [BrandRsp](docs/BrandRsp.md)
 - [BrandRspBrandInfo](docs/BrandRspBrandInfo.md)
 - [BusiLicense](docs/BusiLicense.md)
 - [CategoryGetRsp](docs/CategoryGetRsp.md)
 - [CategoryRsp](docs/CategoryRsp.md)
 - [Cats](docs/Cats.md)
 - [CheckAuditStatusRsp](docs/CheckAuditStatusRsp.md)
 - [CheckAuditStatusRspData](docs/CheckAuditStatusRspData.md)
 - [CheckAuditStatusRspDataPayAuditDetail](docs/CheckAuditStatusRspDataPayAuditDetail.md)
 - [Coupon](docs/Coupon.md)
 - [CouponCouponInfo](docs/CouponCouponInfo.md)
 - [CouponCouponInfoValidInfo](docs/CouponCouponInfoValidInfo.md)
 - [CouponGetListRsp](docs/CouponGetListRsp.md)
 - [CouponInfo](docs/CouponInfo.md)
 - [CouponInfoDiscountInfo](docs/CouponInfoDiscountInfo.md)
 - [CouponInfoDiscountInfoDiscountCondition](docs/CouponInfoDiscountInfoDiscountCondition.md)
 - [CouponInfoExtInfo](docs/CouponInfoExtInfo.md)
 - [CouponInfoPromoteInfo](docs/CouponInfoPromoteInfo.md)
 - [CouponInfoReceiveInfo](docs/CouponInfoReceiveInfo.md)
 - [CouponInfoValidInfo](docs/CouponInfoValidInfo.md)
 - [CouponPushRsp](docs/CouponPushRsp.md)
 - [CouponStockInfo](docs/CouponStockInfo.md)
 - [DeliveryGetCompanyListRsp](docs/DeliveryGetCompanyListRsp.md)
 - [DeliveryGetCompanyListRspData](docs/DeliveryGetCompanyListRspData.md)
 - [DeliveryGetCompanyListRspDataCompanyList](docs/DeliveryGetCompanyListRspDataCompanyList.md)
 - [DeliveryGetFreightTemplateRsp](docs/DeliveryGetFreightTemplateRsp.md)
 - [DeliveryInfo](docs/DeliveryInfo.md)
 - [DeliveryInfoDeliveryProductInfo](docs/DeliveryInfoDeliveryProductInfo.md)
 - [DeliverySendRsp](docs/DeliverySendRsp.md)
 - [FreightTemplateRsp](docs/FreightTemplateRsp.md)
 - [IdCardInfo](docs/IdCardInfo.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject10](docs/InlineObject10.md)
 - [InlineObject11](docs/InlineObject11.md)
 - [InlineObject12](docs/InlineObject12.md)
 - [InlineObject13](docs/InlineObject13.md)
 - [InlineObject14](docs/InlineObject14.md)
 - [InlineObject15](docs/InlineObject15.md)
 - [InlineObject16](docs/InlineObject16.md)
 - [InlineObject17](docs/InlineObject17.md)
 - [InlineObject18](docs/InlineObject18.md)
 - [InlineObject19](docs/InlineObject19.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject20](docs/InlineObject20.md)
 - [InlineObject21](docs/InlineObject21.md)
 - [InlineObject22](docs/InlineObject22.md)
 - [InlineObject23](docs/InlineObject23.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [InlineObject5](docs/InlineObject5.md)
 - [InlineObject6](docs/InlineObject6.md)
 - [InlineObject7](docs/InlineObject7.md)
 - [InlineObject8](docs/InlineObject8.md)
 - [InlineObject9](docs/InlineObject9.md)
 - [NameInfo](docs/NameInfo.md)
 - [Order](docs/Order.md)
 - [OrderExtInfo](docs/OrderExtInfo.md)
 - [OrderGetListRsp](docs/OrderGetListRsp.md)
 - [OrderGetRsp](docs/OrderGetRsp.md)
 - [OrderOrderDetail](docs/OrderOrderDetail.md)
 - [OrderSearchRsp](docs/OrderSearchRsp.md)
 - [OrganizationCodeInfo](docs/OrganizationCodeInfo.md)
 - [PayInfo](docs/PayInfo.md)
 - [PicFile](docs/PicFile.md)
 - [PproductInfo](docs/PproductInfo.md)
 - [PriceInfo](docs/PriceInfo.md)
 - [ProductDeliverySendDeliveryList](docs/ProductDeliverySendDeliveryList.md)
 - [RegisterShopRsp](docs/RegisterShopRsp.md)
 - [ReturnInfo](docs/ReturnInfo.md)
 - [ShopcatRsp](docs/ShopcatRsp.md)
 - [Sku](docs/Sku.md)
 - [SkuAddRsp](docs/SkuAddRsp.md)
 - [SkuAddRspData](docs/SkuAddRspData.md)
 - [SkuBatchAddReq](docs/SkuBatchAddReq.md)
 - [SkuBatchAddRsp](docs/SkuBatchAddRsp.md)
 - [SkuDelRsp](docs/SkuDelRsp.md)
 - [SkuGetListRsp](docs/SkuGetListRsp.md)
 - [SkuGetRsp](docs/SkuGetRsp.md)
 - [SkuUpdatePriceRsp](docs/SkuUpdatePriceRsp.md)
 - [SkuUpdateReq](docs/SkuUpdateReq.md)
 - [SkuUpdateRsp](docs/SkuUpdateRsp.md)
 - [SkuUpdateRspData](docs/SkuUpdateRspData.md)
 - [Spu](docs/Spu.md)
 - [SpuAddReq](docs/SpuAddReq.md)
 - [SpuAddReqDescInfo](docs/SpuAddReqDescInfo.md)
 - [SpuAddReqExpressInfo](docs/SpuAddReqExpressInfo.md)
 - [SpuAddRsp](docs/SpuAddRsp.md)
 - [SpuAddRspData](docs/SpuAddRspData.md)
 - [SpuDelRsp](docs/SpuDelRsp.md)
 - [SpuDelistingRsp](docs/SpuDelistingRsp.md)
 - [SpuDescInfo](docs/SpuDescInfo.md)
 - [SpuGetListRsp](docs/SpuGetListRsp.md)
 - [SpuGetRsp](docs/SpuGetRsp.md)
 - [SpuGetRspData](docs/SpuGetRspData.md)
 - [SpuListingRsp](docs/SpuListingRsp.md)
 - [SpuSearchRsp](docs/SpuSearchRsp.md)
 - [SpuShopcat](docs/SpuShopcat.md)
 - [SpuSkus](docs/SpuSkus.md)
 - [SpuUpdateReq](docs/SpuUpdateReq.md)
 - [SpuUpdateReqExpressInfo](docs/SpuUpdateReqExpressInfo.md)
 - [SpuUpdateRsp](docs/SpuUpdateRsp.md)
 - [SpuUpdateRspData](docs/SpuUpdateRspData.md)
 - [StockGetRsp](docs/StockGetRsp.md)
 - [StockGetRspData](docs/StockGetRspData.md)
 - [StockUpdateRsp](docs/StockUpdateRsp.md)
 - [StockUpdateRspData](docs/StockUpdateRspData.md)
 - [StoreGetInfoRsp](docs/StoreGetInfoRsp.md)
 - [StoreGetInfoRspData](docs/StoreGetInfoRspData.md)
 - [StoreGetShopcatRsp](docs/StoreGetShopcatRsp.md)
 - [SubmitBasicinfo](docs/SubmitBasicinfo.md)
 - [SubmitMerchantinfo](docs/SubmitMerchantinfo.md)
 - [SuperAdministratorInfo](docs/SuperAdministratorInfo.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



