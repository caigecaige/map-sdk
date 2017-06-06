// 高德地图
package sdk

import (
	"encoding/json"
)

type Amap struct {
	Key              string
	Version          string
	BaseUrl          string `ms:"http://restapi.amap.com"`
	ResponseDataType string `ms:"json"`
}

const (
	AMAP_API_METHOD_URI_STATIC_MAP    = "/v3/staticmap"
	AMAP_API_METHOD_URI_CONVERT_COORD = "/v3/assistant/coordinate/convert"
)

func (m *Amap) ApiKey(key string) {
	m.Key = key
}

func (m Amap) MakeStaticMapImage(params map[string]string) string {
	params["key"] = m.Key
	targetUrl := getTagVal(m, "BaseUrl") + AMAP_API_METHOD_URI_STATIC_MAP + "?" + parseParams(params)
	return targetUrl
}

func (m Amap) ConvertCoord(params map[string]string) string {
	params["key"] = m.Key
	targetUrl := getTagVal(m, "BaseUrl") + AMAP_API_METHOD_URI_CONVERT_COORD + "?" + parseParams(params)
	resp := get(targetUrl)
	respJson := make(map[string]string)
	jsonErr := json.Unmarshal(resp, &respJson)
	if jsonErr != nil {
		return ""
	}
	return respJson["locations"]
}

func (m Amap) LocationWithIP(params map[string]string) map[string]interface{} {
	return map[string]interface{}{}
}

func (m Amap) CalculateDistance(params map[string]string) map[string]interface{} {
	return map[string]interface{}{}
}
