package util

import (
	"encoding/json"
	"fmt"
)

func ToJSONIgnoreError(v interface{}, pretty ...bool) string {
	var jsonBytes []byte
	var err error

	if len(pretty) > 0 && pretty[0] {
		jsonBytes, err = json.MarshalIndent(v, "", " ")
	} else {
		jsonBytes, err = json.Marshal(v)
	}
	if err != nil {
		return fmt.Sprintf("error:%s", err)
	}
	return string(jsonBytes)
}

// 将对象转成json字符串
func ToJSON(v interface{}) string {
	var jsonBytes []byte
	var err error

	if v == nil {
		return "{}"
	}

	jsonBytes, err = json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("error:%s", err)
	}
	return string(jsonBytes)
}

// 将json字符串转成目标类型。
func ParseJSON(value string, target interface{}) (interface{}, error) {
	if value == "{}" {
		return nil, nil
	}
	err := json.Unmarshal([]byte(value), target)
	return target, err
}

// 将json字符串转成int64数组。
func ParseJSONToInt64Arr(json string) ([]int64, error) {
	var arr []int64
	if json == "" || json == "[]" {
		return arr, nil
	}
	_, err := ParseJSON(json, &arr)
	if err != nil {
		return nil, err
	}

	return arr, nil
}

// 将json字符串转成string数组。
func ParseJSONToStringArr(json string) ([]string, error) {
	var arr []string
	if json == "" || json == "[]" {
		return arr, nil
	}
	_, err := ParseJSON(json, &arr)
	if err != nil {
		return nil, err
	}

	return arr, nil
}
