package z

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 将一个字节数组转换成 utf-8 字符串
func Utf8(bs []byte) (str string, err error) {
	if utf8.FullRune(bs) {
		//sz := utf8.RuneCount(bs)
		str = string(bs)
		return
	}
	// 错误
	err = errors.New("fail to decode to UTF8")
	str = ""
	return
}

// 是不是空字符
func IsSpace(c byte) bool {
	if c >= 0x00 && c <= 0x20 {
		return true
	}
	return false
}

// 判断一个字符串是不是空白串，即（0x00 - 0x20 之内的字符均为空白字符）
func IsBlank(s string) bool {
	for i := 0; i < len(s); i++ {
		b := s[i]
		if !IsSpace(b) {
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

// 去掉多余的字符串
// 比如 " a b  c    d e" -> "a b c d e"
func TrimExtraSpace(s string) string {
	s = Trim(s)
	size := len(s)
	sb := SBuilder()
	isSpace := false
	for i := 0; i < size; i++ {
		c := s[i]
		if !IsSpace(c) {
			if isSpace {
				sb.AppendByte(' ')
				isSpace = false
			}
			sb.AppendByte(c)
		} else {
			if !isSpace {
				isSpace = true
			}
		}
	}
	return sb.String()
}

// 复制字符
func DupChar(char byte, num int) string {
	sb := SBuilder()
	for i := 0; i < num; i++ {
		sb.AppendByte(char)
	}
	return sb.String()
}

// 复制字符串
func Dup(str string, num int) string {
	sb := SBuilder()
	for i := 0; i < num; i++ {
		sb.Append(str)
	}
	return sb.String()
}

// 填充字符串右侧一定数量的特殊字符
func AlignLeft(str string, width int, char byte) string {
	length := len(str)
	if length < width {
		return str + DupChar(char, width-length)
	}
	return str
}

// 填充字符串左侧一定数量的特殊字符
func AlignRight(str string, width int, char byte) string {
	length := len(str)
	if length < width {
		return DupChar(char, width-length) + str
	}
	return str
}

type stringBuilder struct {
	buf *bytes.Buffer
}

// 提供一个类似java中stringBuilder对象,支持链式调用(不返回错误信息)
func SBuilder() *stringBuilder {
	sb := new(stringBuilder)
	sb.buf = bytes.NewBuffer(nil)
	return sb
}

// 添加字符串
func (sb *stringBuilder) Append(str string) *stringBuilder {
	_, err := sb.buf.WriteString(str)
	if err != nil {
		panic(err)
	}
	return sb
}

// 添加字符
func (sb *stringBuilder) AppendByte(char byte) *stringBuilder {
	err := sb.buf.WriteByte(char)
	if err != nil {
		panic(err)
	}
	return sb
}

// 添加字符数组, 会自动添加"[]"并用","做分割
func (sb *stringBuilder) AppendByteArray(chars []byte) *stringBuilder {
	sb.AppendByte('[')
	for i, char := range chars {
		sb.AppendByte('\'').AppendByte(char).AppendByte('\'')
		if i != len(chars)-1 {
			sb.Append(", ")
		}
	}
	sb.AppendByte(']')
	return sb
}

// 添加字符串数组, 会自动添加"[]"并用","做分割
func (sb *stringBuilder) AppendStringArray(strs []string) *stringBuilder {
	sb.AppendByte('[')
	for i, str := range strs {
		sb.AppendByte('\'').Append(str).AppendByte('\'')
		if i != len(strs)-1 {
			sb.Append(", ")
		}
	}
	sb.AppendByte(']')
	return sb
}

// 返回字符串
func (sb *stringBuilder) String() string {
	return sb.buf.String()
}
