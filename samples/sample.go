// sample.go
package main

import (
	"fmt"
	"map-sdk"
)

func main() {
	ms := mapsdk.New("amap", "your api key")
	params := make(map[string]string)

	params["locations"] = "114.028542,22.616764"
	params["coordsys"] = "gps"
	res := ms.ConvertCoord(params)
	fmt.Println(res)
}
