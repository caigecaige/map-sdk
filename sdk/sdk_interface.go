// sdk_interface
package sdk

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

const STRUCT_TAG = "ms"

type SdkInterface interface {
	ApiKey(string)
	MakeStaticMapImage(map[string]string) string                //静态图
	ConvertCoord(map[string]string) string                      //转换坐标
	LocationWithIP(map[string]string) map[string]interface{}    //ip定位
	CalculateDistance(map[string]string) map[string]interface{} //计算距离
}

//通用的struct
type Coord struct {
	Longitude float64 `json:"lng"`
	Latitude  float64 `json:"lat"`
}

//通用方法
func getTagVal(i interface{}, field string) string {
	structField, getErr := reflect.TypeOf(i).FieldByName(field)
	if !getErr {
		return ""
	}
	tag := structField.Tag.Get(STRUCT_TAG)
	return tag
}

func parseParams(params map[string]string) string {
	paramStr := ""
	if len(params) > 0 {
		for k, v := range params {
			paramStr = paramStr + "&" + k + "=" + v
		}
		paramStr = strings.Trim(paramStr, "&")
	}
	return paramStr
}

func post() {

}

func get(url string) []byte {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil
	}
	respContent, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil
	}
	return respContent
}

func float2string(f float64) string {
	return strconv.FormatFloat(f, 'f', 7, 64)
}
