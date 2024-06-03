package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eyebluecn/sc-misc/src/util/uuid"
	"regexp"
	"strconv"
	"strings"
)

// StringToInt64 string => int64
func StringToInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func StringsToInt64s(vals []string) []int64 {
	res := make([]int64, 0)
	for _, val := range vals {
		v, _ := strconv.ParseInt(val, 10, 64)
		res = append(res, v)
	}
	return res
}

func StringToInt32(s string) int32 {
	v, _ := strconv.ParseInt(s, 10, 32)
	return int32(v)
}

func StringToInt8(s string) int8 {
	v, _ := strconv.ParseInt(s, 10, 8)
	return int8(v)
}

func StringToInt64Addr(s string) (*int64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	return &v, err
}

func StringAddrToInt64Addr(sp *string) *int64 {
	if sp == nil {
		return nil
	}
	v, _ := strconv.ParseInt(*sp, 10, 64)
	return &v
}
func StringAddrToInt64(sp *string, defaultVal int64) int64 {
	if sp != nil {
		if v, err := strconv.ParseInt(*sp, 10, 64); err != nil {
			return v
		}
	}
	return defaultVal
}

func StringToInt32Addr(s string) (*int32, error) {
	v, err := strconv.ParseInt(s, 10, 32)
	v1 := int32(v)
	return &v1, err
}

func StringToInt8Addr(s string) (*int8, error) {
	v, err := strconv.ParseInt(s, 10, 8)
	v1 := int8(v)
	return &v1, err
}

func Int64ToString(num *int64) string {
	if num == nil {
		return ""
	}
	return strconv.FormatInt(*num, 10)
}

func Int64ToString2(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Int64ValueToString(num int64) string {
	v := strconv.FormatInt(num, 10)
	return v
}

func Int64ToStringAddr(num int64) *string {
	v := strconv.FormatInt(num, 10)
	return &v
}

func StringAddrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func Int64AddrToStringAddr(num *int64) *string {
	if num == nil {
		return nil
	}
	v := strconv.FormatInt(*num, 10)
	return &v
}

func Int32ToString(num *int32) string {
	if num == nil {
		return ""
	}
	return strconv.FormatInt(int64(*num), 10)
}

func Int8ToString(num *int8) string {
	if num == nil {
		return ""
	}
	return strconv.FormatInt(int64(*num), 10)
}

func Float64ToString(num *float64) string {
	if num == nil {
		return ""
	}
	return strconv.FormatFloat(*num, 'f', -1, 64)
}

func StructToStr(data interface{}) string {
	d, _ := json.Marshal(data)
	return string(d)
}

func NewInt32(val int32) *int32 {
	ret := val
	return &ret
}

func NewInt8(val int8) *int8 {
	ret := val
	return &ret
}

func NewInt64(val int64) *int64 {
	ret := val
	return &ret
}

func NewString(val string) *string {
	ret := val
	return &ret
}

func Print(data interface{}, str ...string) {
	d, _ := json.Marshal(data)
	fmt.Println(str, "---------------------"+string(d))
}

// 去掉开头的引号 excel使用
func TrimExcelQuotes(s string) string {
	if s != "" {
		reg := regexp.MustCompile("[\\r\\n\\t]")
		s = strings.Trim(reg.ReplaceAllString(s, ""), "' ")
	}
	return s
}

func StringToPtr(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func ToStringPtr(num *int64) *string {
	if num == nil {
		return nil
	}
	str := strconv.FormatInt(*num, 10)
	return &str
}

func StringToSlice(str *string) []string {
	slice := make([]string, 0)
	if str == nil {
		return slice
	}
	return strings.Split(*str, `,`)
}

func StringToSliceInt32(str *string) []int32 {
	slice := make([]int32, 0)
	if str == nil {
		return slice
	}
	strArr := strings.Split(*str, `,`)
	if len(strArr) <= 0 {
		return slice
	}
	for _, s := range strArr {
		num := StringToInt32(s)
		slice = append(slice, num)
	}
	return slice
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// 去掉前面和后面的空格
func Trim(s string) string {

	return strings.Trim(s, " ")
}

func IsEmpty(str *string) bool {
	return str == nil || len(*str) == 0
}

func SliceLast(s, sep string) string {
	split := strings.Split(s, sep)
	if len(split) == 0 {
		return ""
	}
	return split[len(split)-1]
}

func StringArr2InterfaceArr(itemList []string) []interface{} {
	res := make([]interface{}, len(itemList))
	for i, item := range itemList {
		res[i] = item
	}
	return res
}
func Int64Arr2InterfaceArr(itemList []int64) []interface{} {
	res := make([]interface{}, len(itemList))
	for i, item := range itemList {
		res[i] = item
	}
	return res
}
func IsEmptyString(s *string) bool {
	if s == nil {
		return true
	}
	if *s == "" {
		return true
	}
	return false
}

type StringBuilder struct {
	buf bytes.Buffer
}

func NewStringBuilder(s ...string) *StringBuilder {
	sb := StringBuilder{}
	sb.buf.WriteString(strings.Join(s, ""))
	return &sb
}

func (sb *StringBuilder) Append(s string) *StringBuilder {
	sb.buf.WriteString(s)
	return sb
}

func (sb *StringBuilder) AppendFormat(f string, s ...interface{}) *StringBuilder {
	sb.buf.WriteString(fmt.Sprintf(f, s...))
	return sb
}

func (sb *StringBuilder) String() string {
	return sb.ToString()
}

func (sb *StringBuilder) ToString() string {
	return sb.buf.String()
}

func StartWith(s string, prefix string) bool {
	if prefix == "" {
		return true
	}
	if s == "" {
		return false
	}
	if len(s) < len(prefix) {
		return false
	}
	return s[:len(prefix)] == prefix
}

func ReplaceAll(str string, old string, new string) string {
	return strings.Replace(str, old, new, -1)
}

// 替换掉前缀的一些内容
func ReplacePrefix(str string, old string, new string) string {
	if StartWith(str, old) {
		return new + str[len(old):]
	} else {
		return str
	}
}

// 生成一个uuid
func Uuid() string {
	timeUUID, _ := uuid.NewV4()
	return timeUUID.String()
}
