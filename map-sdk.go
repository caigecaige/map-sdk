// sdk调用，新加入的sdk要先注册
package mapsdk

import (
	"map-sdk/sdk"
)

type MapSdk struct {
	Name          string
	RegisteredSdk map[string]sdk.SdkInterface
	Sdk           sdk.SdkInterface
}

func New(name, key string) *MapSdk {
	ms := new(MapSdk)
	ms.registerSdk()
	ms.Sdk = ms.RegisteredSdk[name]
	ms.Sdk.ApiKey(key)
	return ms
}

//注册sdk
func (ms *MapSdk) registerSdk() {
	ms.RegisteredSdk = make(map[string]sdk.SdkInterface)
	ms.RegisteredSdk["qq"] = new(sdk.QQMap)
	ms.RegisteredSdk["amap"] = new(sdk.Amap)
}

//静态地图
func (ms *MapSdk) MakeStaticMapImage(params map[string]string) string {
	return ms.Sdk.MakeStaticMapImage(params)
}

//转换坐标
func (ms *MapSdk) ConvertCoord(params map[string]string) string {
	return ms.Sdk.ConvertCoord(params)
}

//ip定位
func (ms *MapSdk) LocationWithIP(params map[string]string) map[string]interface{} {
	return ms.Sdk.LocationWithIP(params)
}

//计算距离
func (ms *MapSdk) CalculateDistance(params map[string]string) map[string]interface{} {
	return ms.Sdk.CalculateDistance(params)
}
