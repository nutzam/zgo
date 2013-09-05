package z

import (
	"strconv"
	"strings"
)

// 判断一个字符串是不是空白串，即（0x00 - 0x20 之内的字符均为空白字符）
func IsBlank(s string) bool {
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b < 0 || b > 0x20 {
			return false
		}
	}
	return true
}

// 如果输入的字符串为空串，那么返回默认字符串
func SBlank(s, dft string) string {
	if IsBlank(s) {
		return dft
	}
	return s
}

// 将字符串转换成整数，如果转换失败，采用默认值
func ToInt(s string, dft int) int {
	var re, err = strconv.Atoi(s)
	if err != nil {
		return dft
	}
	return re
}

// 将字符串转换成整数，如果转换失败，采用默认值
func ToInt64(s string, dft int64) int64 {
	var re, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return dft
	}
	return re
}

// 拆分字符串数组，如果数组元素为空白，忽略，否则则 Trim 空白
func SplitIgnoreBlank(s, sep string) []string {
	ss := strings.Split(s, sep)
	size := len(ss)
	re := make([]string, 0, size)
	for i := 0; i < size; i++ {
		str := Trim(ss[i])
		//log.Printf("%d : '%s'", i, str)
		if len(str) > 0 {
			//log.Printf("  append @ [%d]", len(re))
			re = append(re, str)
		}
	}
	return re
}

// 去掉一个字符串左右的空白串，即（0x00 - 0x20 之内的字符均为空白字符）
func Trim(s string) string {
	size := len(s)
	if size <= 0 {
		return s
	}
	l := 0
	for ; l < size; l++ {
		b := s[l]
		if b < 0 || b > 0x20 {
			//log.Printf("l stop %d : '%c'", l, b)
			break
		}
	}
	r := size - 1
	for ; r >= l; r-- {
		b := s[r]
		if b < 0 || b > 0x20 {
			//log.Printf("r stop %d : '%c'", r, b)
			break
		}
	}
	return string(s[l : r+1])
}

// 去掉一个字符串左右的空白串，即（0x00 - 0x20 之内的字符均为空白字符）
func TrimBytes(bs []byte) string {
	r := len(bs) - 1
	if r <= 0 {
		return string(bs)
	}
	l := 0
	for ; l <= r; l++ {
		b := bs[l]
		if b < 0 || b > 0x20 {
			break
		}
	}
	for ; r >= l; r-- {
		b := bs[r]
		if b < 0 || b > 0x20 {
			break
		}
	}
	return string(bs[l : r+1])
}
