package util

import (
	"net/url"
	"os"
	"strings"
)

// 判断一个路径是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return false
		}
	} else {
		return true
	}
}

// 从一个query中解析出对应参数。 可以返回数组，出错了会直接报错
func ParseQuery(query string, key string) ([]string, error) {
	u, err := url.Parse(query)
	if err != nil {
		return nil, err
	}

	values, ok := u.Query()[key]
	if !ok {
		return nil, err
	}

	return values, nil
}

// 从一个query中解析出对应参数。 如果出错了，或者不存在，均返回空字符串
func SimpleParseQuery(query string, key string) string {

	values, err := ParseQuery(query, key)
	if err != nil {
		return ""
	}

	if len(values) > 0 {
		return values[0]
	}

	//尝试添加一个?号后再解析。
	if !strings.Contains(query, "?") {
		values, err = ParseQuery("?"+query, key)
	}

	if len(values) > 0 {
		return values[0]
	}

	return ""
}
