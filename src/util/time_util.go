package util

import "time"

const (
	LAYOUT_yyyy_MM_dd          = "2006-01-02"
	LAYOUT_yyyy_MM_dd_HH_mm_ss = "2006-01-02 15:04:05"
	LAYOUT_yyyy_MM             = "2006-01"
	LAYOUT_yyyyMMdd            = "20060102"
	LAYOUT_yyyyMMddHHmm        = "200601021504"
	LAYOUT_yyyyMMddHHmmss      = "20060102150405"
)

// 时间转为时间戳
func Timestamp(time time.Time) int64 {
	return time.Unix()
}

func YesterdayyyyyMMdd() string {
	nTime := time.Now()
	yesterday := nTime.AddDate(0, 0, -1)
	return FmtTime2Str(yesterday, LAYOUT_yyyyMMdd)
}

func FmtTime2Str(time time.Time, layout string) string {
	formatTime := time.Format(layout)
	return formatTime
}

// 转为 yyyy-MM-dd HH:mm:ss 的字符串格式
func FormatDateTime(time time.Time) string {
	formatTime := time.Format(LAYOUT_yyyy_MM_dd_HH_mm_ss)
	return formatTime
}

func ParseDate(dateStr string) (time.Time, error) {
	t, err := time.ParseInLocation(LAYOUT_yyyy_MM_dd_HH_mm_ss, dateStr, time.Local)
	return t, err
}

// 时间戳转为时间
func ParseTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func ParseDateInLayout(dateStr string, layout string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, dateStr, time.Local)
	return t, err
}

func AddDay(t time.Time, day int32) time.Time {
	return t.Add(time.Duration(day*24) * time.Hour)
}

// 将日期时间字符串转为秒。
func DateTimeToSeconds(dateTimeStr string) int64 {
	t, err := time.ParseInLocation(LAYOUT_yyyy_MM_dd_HH_mm_ss, dateTimeStr, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 将秒数变为日期时间字符串
func SecondsToDateTime(data int64) string {

	tm := time.Unix(data, 0)

	return tm.Format(LAYOUT_yyyy_MM_dd_HH_mm_ss)

}
