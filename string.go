package z

import (
	"bytes"
	"errors"
	"fmt"
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
		if len(str) > 0 {
			re = append(re, str)
		}
	}
	return re
}

// 去掉一个字符串左右的空白串，即（0x00 - 0x20 之内的字符均为空白字符）
// 与strings.TrimSpace功能一致
func Trim(s string) string {
	size := len(s)
	if size <= 0 {
		return s
	}
	l := 0
	for ; l < size; l++ {
		b := s[l]
		if !IsSpace(b) {
			break
		}
	}
	r := size - 1
	for ; r >= l; r-- {
		b := s[r]
		if !IsSpace(b) {
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
		if !IsSpace(b) {
			break
		}
	}
	for ; r >= l; r-- {
		b := bs[r]
		if !IsSpace(b) {
			break
		}
	}
	return string(bs[l : r+1])
}

// Trim并且去掉中间多余的空白(多个空白变一个空白)
// 比如 " a b  c    d e" -> "a b c d e"
func TrimExtraSpace(s string) string {
	s = Trim(s)
	size := len(s)
	switch size {
	case 0, 1, 2, 3:
		return s
	default:
		bs := make([]byte, 0, size)
		isSpace := false
		for i := 0; i < size; i++ {
			c := s[i]
			if !IsSpace(c) {
				if isSpace {
					bs = append(bs, ' ')
					isSpace = false
				}
				bs = append(bs, c)
			} else {
				if !isSpace {
					isSpace = true
				}
			}
		}
		return string(bs)
	}
	// 兼容低版本GO
	return ""
}

// 复制字符
func DupChar(char byte, num int) string {
	bs := make([]byte, num, num)
	for i := 0; i < num; i++ {
		bs[i] = char
	}
	return string(bs)
}

// 复制字符串
func Dup(str string, num int) string {
	return strings.Repeat(str, num)
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

type strBuilder struct {
	buf *bytes.Buffer
}

// 提供一个类似java中stringBuilder对象,支持链式调用(不返回错误信息,直接panic)
// str := SBuilder().Append("abc=123").Append('\n').String()
// TODO 等着测试下性能,看看用字符数组来实现是不是效率高些
func StringBuilder() *strBuilder {
	sb := new(strBuilder)
	sb.buf = bytes.NewBuffer(nil)
	return sb
}

// 添加任意可以生成string的东西
func (sb *strBuilder) Append(o interface{}) *strBuilder {
	var err error
	switch o.(type) {
	case byte:
		b, _ := o.(byte)
		err = sb.buf.WriteByte(b)
	case rune:
		r, _ := o.(rune)
		_, err = sb.buf.WriteRune(r)
	default:
		str := fmt.Sprint(o)
		_, err = sb.buf.WriteString(str)
	}
	if err != nil {
		panic(err)
	}
	return sb
}

// 行结尾了换行(EOL End Of Line)
func (sb *strBuilder) EOL() *strBuilder {
	sb.Append('\n')
	return sb
}

// 返回字符串
func (sb *strBuilder) String() string {
	return sb.buf.String()
}

// 写入的字符串长度
func (sb *strBuilder) Len() int {
	return sb.buf.Len()
}

// 字符串转Float
func ToFloat(data string, reData float64) float64 {
	i, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return reData
	}
	return i

}
