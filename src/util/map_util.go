package util

import "reflect"

func PoToMap(po interface{}, ignoreZeroFields []string) map[string]interface{} {
	if po == nil {
		return nil
	}
	recordType := reflect.TypeOf(po)
	if recordType.Kind() == reflect.Ptr {
		structValue := reflect.ValueOf(po).Elem().Interface()
		return doPoToMap(structValue, ignoreZeroFields)
	}
	return doPoToMap(po, ignoreZeroFields)
}

func doPoToMap(po interface{}, ignoreZeroFields []string) map[string]interface{} {
	// json和copier都不可用
	// json会把所有的数字转成float64
	// copier不支持struct转map
	result := make(map[string]interface{})

	ignoreMap := make(map[string]bool, 0)
	for _, item := range ignoreZeroFields {
		ignoreMap[item] = true
	}

	recordType := reflect.TypeOf(po)
	for i := 0; i < recordType.NumField(); i++ {
		field := recordType.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}

		fieldValue := reflect.ValueOf(po).FieldByIndex(field.Index)
		if fieldValue.IsZero() && ignoreMap[field.Name] == true {
			// 跳过不需要0值的属性
			continue
		}
		result[jsonTag] = fieldValue.Interface()
	}

	return result
}
