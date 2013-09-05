package z

// 本方法确保没错错误，如果有错误，则强制退出整个程序
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}
