package pkg

import "github.com/xgpc/util"

func StructToMap(obj interface{}) map[string]interface{} {
	body, err := util.JsonEncode(obj)
	if err != nil {
		panic(err)
	}
	var md map[string]interface{}
	err = util.JsonDecode(body, &md)
	if err != nil {
		panic(err)
	}
	return md
}
