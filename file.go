package z

import (
	"errors"
	"log"
	"os"
	"path"
)

// Remove 文件
func Fremove(name string) (err error) {
	err = os.Remove(name)
	return err
}

// 创建一个空文件，如果文件已存在，返回 false
func Fnew(ph string) error {
	if Exists(ph) {
		return errors.New("file does not exist" + " " + ph)
	}
	// 确保父目录存在
	err := Mkdir(path.Dir(ph))
	if err != nil {
		return err
	}
	// 创建
	_, err = os.Create(ph)
	if nil != err {
		return err
	}
	return nil
}

/*
调用者将负责关闭文件
*/
func FileA(ph string) *os.File {
	// 确定文件的父目录是存在的
	FcheckParents(ph)
	// 打开文件，文件不存在则创建,追加方式
	f, err := os.OpenFile(ph, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if nil != err {
		panic(err)
	}
	return f
}

// 用回调的方式打文件以便追加内容，回调函数不需要关心文件关闭等问题
func FileAF(ph string, callback func(*os.File)) {
	f := FileA(ph)
	if nil != f {
		defer f.Close()
		callback(f)
	}
}

// 打开一个文件准备复写内容，如果文件不存在，则创建它
// 如果有错误，将打印 log
//
// 调用者将负责关闭文件
func FileW(ph string) *os.File {
	// 确定文件的父目录是存在的
	FcheckParents(ph)
	// 打开文件
	f, err := os.OpenFile(ph, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if nil != err {
		panic(err)
	}
	return f
}

// 用回调的方式打文件以便复写内容，回调函数不需要关心文件关闭等问题
func FileWF(ph string, callback func(*os.File)) {
	f := FileW(ph)
	// 打开失败，那么将试图创建
	if nil == f && !Fexists(ph) {
		err := Fnew(ph)
		if err != nil {
			panic(err)
		}
		f = FileW(ph)
	}
	// 开始写入
	if nil != f {
		defer f.Close()
		if nil != callback {
			callback(f)
		}
	}
}

/*
将从自己磁盘目录，只读的方式打开一个文件。如果文件不存在，或者打开错误，则返回 nil。
如果有错误，将打印 log

调用者将负责关闭文件
*/
func FileR(ph string) *os.File {
	f, err := os.Open(ph)
	if nil != err {
		return nil
	}
	return f
}

// 用回调的方式打文件以便读取内容，回调函数不需要关心文件关闭等问题
func FileRF(ph string, callback func(*os.File)) {
	f := FileR(ph)
	if nil != f {
		defer f.Close()
		callback(f)
	}
}

// 自定义模式打开文件
// 如果有错误，将打印 log 并返回 nil
// 调用者将负责关闭文件
func FileO(ph string, flag int) *os.File {
	// 确定文件的父目录是存在的
	FcheckParents(ph)
	// 打开文件
	f, err := os.OpenFile(ph, flag, 0666)
	if nil != err {
		log.Println(err)
		return nil
	}
	return f
}

// 用自定义的模式打文件以便替换内容，回调函数不需要关心文件关闭等问题
func FileOF(ph string, flag int, callback func(*os.File)) {
	f := FileO(ph, flag)
	// 开始写入
	if nil != f {
		defer f.Close()
		if nil != callback {
			callback(f)
		}
	}
}
