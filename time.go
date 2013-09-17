package z

import (
	"strconv"
	"time"
)

const FORMAT_DATE string = "2006-01-02"
const FORMAT_DATE_TIME string = "2006-01-02 15:04:05"

// 获取本地时间戳纳秒,以字符串格式返回
func UnixNano() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// 获取本地事件毫秒
func UnixMsSec(off int) int64 {
	return (time.Now().Unix() + int64(off)) * 1000
}

// 获取某个时间当天的绝对秒数
func DAoffSec(t time.Time, off int) int {
	hour, min, sec := t.Clock()
	return hour*3600 + min*60 + sec + off
}

// 获取当天的绝对秒数，根据当前时间 +n
func DAsec(off int) int {
	return DAoffSec(time.Now(), off)
}

// 获得当前系统时间
func GetTime() string {
	return time.Now().Format(FORMAT_DATE_TIME)
}

func ParseDate(dstr string) time.Time {
	t, _ := time.Parse(FORMAT_DATE, dstr)
	return t
}

func ParseDateTime(dtstr string) time.Time {
	t, _ := time.Parse(FORMAT_DATE_TIME, dtstr)
	return t
}
