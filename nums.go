package z

// 返回byte(char)在数组中的位置
func IndexOfBytes(array []byte, one byte) int {
	for i, entity := range array {
		if entity == one {
			return i
		}
	}
	return -1
}

// 返回字符串在数组中的位置
func IndexOfStrings(array []string, one string) int {
	for i, entity := range array {
		if entity == one {
			return i
		}
	}
	return -1
}

// 判断字符串是否在数组中
func IsInStrings(array []string, one string) bool {
	return IndexOfStrings(array, one) > -1
}
