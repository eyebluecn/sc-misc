package util

import (
	"net/url"
)

func SetQueryParam(input string, key string, value string) (string, error) {
	if len(input) == 0 || len(key) == 0 || len(value) == 0 {
		return input, nil
	}
	u, err := url.Parse(input)
	if err != nil {
		return input, err
	}
	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return input, err
	}
	query.Set(key, value)
	u.RawQuery = query.Encode()
	return u.String(), nil
}

// SetSchemaURLQueryParam Android H5页面不能获取schema 上面的query param，所以需要先解析url，把query param 加到url 上
func SetSchemaURLQueryParam(schema string, key string, value string) (string, error) {
	if len(schema) == 0 || len(key) == 0 || len(value) == 0 {
		return schema, nil
	}
	u, err := url.Parse(schema)
	if err != nil {
		return schema, err
	}
	schemaQuery, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return schema, err
	}
	schemaURL := schemaQuery.Get("url")
	if len(schemaURL) == 0 {
		return schema, nil
	}
	newSchemaURL, err := SetQueryParam(schemaURL, key, value)
	if err != nil {
		return schema, err
	}
	if len(newSchemaURL) == 0 {
		return schema, nil
	}
	schemaQuery.Set("url", newSchemaURL)
	u.RawQuery = schemaQuery.Encode()
	return u.String(), nil
}

// SetSchemaQueryParam 屏蔽lynx/H5 URL细节，在schema和url 上都拼接参数
func SetSchemaQueryParam(schema string, key string, value string) (string, error) {
	if len(schema) == 0 || len(key) == 0 || len(value) == 0 {
		return schema, nil
	}
	schema, _ = SetQueryParam(schema, key, value)
	schema, _ = SetSchemaURLQueryParam(schema, key, value)
	return schema, nil
}
