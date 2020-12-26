/*
 * 微信小商店开放API
 *
 * wei xin xiao shang dian open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wxxiaodian
// Cats 类目
type Cats struct {
	// 类目ID，如果brand_id=2100000000，需要先通过获取类目接口category/get拿到可用的cat_id；如果brand_id!=2100000000，则这里的cat_id需要与获取品牌接口brand/get中的1,2,3级类目一一对应
	CatId int64 `json:"cat_id"`
	// 类目层级
	Level int32 `json:"level"`
}
