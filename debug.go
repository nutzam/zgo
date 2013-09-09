package z

import (
	"fmt"
)

// 给一个总开关, 这样在调试的时候就可以使用总开关关闭打印语句了
var _print_on_ bool = false

func IsDebugOn() bool {
	return _print_on_
}

func DebugOn() {
	_print_on_ = true
}

func DebugOff() {
	_print_on_ = false
}

func DebugPrint(a ...interface{}) (n int, err error) {
	if _print_on_ {
		return fmt.Print(a...)
	}
	return 0, nil
}

func DebugPrintf(format string, a ...interface{}) (n int, err error) {
	if _print_on_ {
		return fmt.Printf(format, a...)
	}
	return 0, nil
}

func DebugPrintln(a ...interface{}) (n int, err error) {
	if _print_on_ {
		return fmt.Println(a...)
	}
	return 0, nil
}
