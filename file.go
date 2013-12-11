package z

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// 按照 UTF8 格式化，将一个磁盘上的文件读成字符串
func Utf8f(ph string) (str string, err error) {
	ph = Ph(ph)
	f, err := os.Open(ph)
	if nil != err {
		str = ""
		return
	}
	str, err = Utf8r(f)
	return
}

// 按照 UTF8 格式，将流的内容读成字符串
func Utf8r(r io.Reader) (str string, err error) {
	bs, err := ioutil.ReadAll(r)
	if nil != err {
		str = ""
		return
	}
	str, err = Utf8(bs)
	return
}

// 便利的获得 FileInfo 对象的方法
func Fi(ph string) os.FileInfo {
	ph = Ph(ph)
	f, e := os.Open(ph)
	NoError(e)
	return Fif(f)
}

// 便利的获得 FileInfo 对象的方法
func Fif(f *os.File) os.FileInfo {
	fi, e := f.Stat()
	NoError(e)
	return fi
}

// 便利的获得文件大小的方法
func Fszf(f *os.File) int64 {
	fi, e := f.Stat()
	NoError(e)
	return fi.Size()
}

// 便利的获得文件大小的方法
func Fsz(ph string) int64 {
	ph = Ph(ph)
	f, e := os.Open(ph)
	NoError(e)
	return Fszf(f)
}

// Remove 文件
func Fremove(ph string) (err error) {
	ph = Ph(ph)
	err = os.Remove(ph)
	return err
}

// 创建一个空文件，如果文件已存在，返回 false
func Fnew(ph string) error {
	ph = Ph(ph)
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
	ph = Ph(ph)
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
	ph = Ph(ph)
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
	ph = Ph(ph)
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
	ph = Ph(ph)
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
	ph = Ph(ph)
	f, err := os.Open(ph)
	if nil != err {
		return nil
	}
	return f
}

// 用回调的方式打文件以便读取内容，回调函数不需要关心文件关闭等问题
func FileRF(ph string, callback func(*os.File)) {
	ph = Ph(ph)
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
	ph = Ph(ph)
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
	ph = Ph(ph)
	f := FileO(ph, flag)
	// 开始写入
	if nil != f {
		defer f.Close()
		if nil != callback {
			callback(f)
		}
	}
}

// 强制覆盖写入文件
func FWrite(path string, data []byte) error {
	// 保证目录存在
	FcheckParents(path)
	// 写入文件
	return ioutil.WriteFile(path, data, 0644)
}
