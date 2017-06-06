// qq地图
package sdk

import (
	"encoding/json"
	"strings"
)

type QQMap struct {
	Key              string
	Version          string
	BaseUrl          string `ms:"http://apis.map.qq.com"`
	ResponseDataType string `ms:"json"`
}

const (
	QQ_MAP_API_METHOD_URI_STATIC_MAP    = "/ws/staticmap/v2/"
	QQ_MAP_API_METHOD_URI_CONVERT_COORD = "/ws/coord/v1/translate"
	//QQ_MAP_API_METHOD_URI_LOCATION_IP   = "/ws/location/v1/ip"
)

func (m *QQMap) ApiKey(key string) {
	m.Key = key
}

func (m QQMap) MakeStaticMapImage(params map[string]string) string {
	params["key"] = m.Key
	targetUrl := getTagVal(m, "BaseUrl") + QQ_MAP_API_METHOD_URI_STATIC_MAP + "?" + parseParams(params)
	return targetUrl
}

const RESULT_CODE_OK = 0

type QQMapConvertCoordResult struct {
	Status    int     `json:"status"`
	Message   string  `json:"message"`
	Locations []Coord `json:"locations"`
}

func (m QQMap) ConvertCoord(params map[string]string) string {
	params["key"] = m.Key
	targetUrl := getTagVal(m, "BaseUrl") + QQ_MAP_API_METHOD_URI_CONVERT_COORD + "?" + parseParams(params)
	resp := get(targetUrl)
	respJson := new(QQMapConvertCoordResult)
	jsonErr := json.Unmarshal(resp, respJson)
	if jsonErr != nil {
		return ""
	}
	if respJson.Status != RESULT_CODE_OK {
		return ""
	}
	coords := ""
	if len(respJson.Locations) > 0 {
		for _, v := range respJson.Locations {
			coords = coords + ";" + float2string(v.Longitude) + "," + float2string(v.Latitude)
		}
	}
	return strings.Trim(coords, ";")
}

func (m QQMap) LocationWithIP(params map[string]string) map[string]interface{} {
	return map[string]interface{}{}
}

func (m QQMap) CalculateDistance(params map[string]string) map[string]interface{} {
	return map[string]interface{}{}
}
