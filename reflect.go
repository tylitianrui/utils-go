package utils

import "reflect"

// 获取某struct的字段的值，如果不存在返回默认值
// 场景，多个struct但都有同一个字段，需要获取此字段的值
func GetStructField(data interface{}, key string, defaultVal interface{}) interface{} {
	typeOfData := reflect.TypeOf(data)
	valueOfData := reflect.ValueOf(data)
	if typeOfData.Kind() == reflect.Struct {
		for i := 0; i < typeOfData.NumField(); i++ {
			if typeOfData.Field(i).Name == key {
				return valueOfData.Field(i).Interface()
			}

		}

	}
	return defaultVal

}
