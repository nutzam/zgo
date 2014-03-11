package z

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 根据给定路径获取绝对路径，可以支持 ~ 作为主目录
func Ph(ph string) string {
	if IsBlank(ph) {
		return ""
	}
	if strings.HasPrefix(ph, "~") {
		home := os.Getenv("HOME")
		if IsBlank(home) {
			panic(fmt.Sprintf("can not found HOME in envs, '%s' AbsPh Failed!", ph))
		}
		ph = fmt.Sprint(home, string(ph[1:]))
	}
	s, err := filepath.Abs(ph)
	if nil != err {
		panic(err)
	}
	return s
}

// 创建一个目录，如果目录存在或者创建成功，返回 true，否则返回 false
func Mkdir(ph string) error {
	err := os.MkdirAll(ph, os.ModeDir|0755)
	if nil != err {
		return err
	}
	return err
}

// 判断一个路径是否存在
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 判断一个路径是否存在,且允许附加条件
func ExistsF(name string, callback func(os.FileInfo) bool) bool {
	fi, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return callback(fi)
}

// 判断一个路径文件是否存在,且不是文件夹
func ExistsIsFile(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	if fi.IsDir() {
		return false
	}
	return true
}

// 判断一个路径文件是否存在,且不是文件夹
func ExistsIsDir(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	if !fi.IsDir() {
		return false
	}
	return true
}

// 判断一个文件是否存在，不存在则创建
func ExistsFile(aph string) bool {
	if Exists(aph) {
		return false
	}
	// 确保父目录存在
	Mkdir(path.Dir(aph))
	// 创建
	_, err := os.Create(aph)
	if nil != err {
		return false
	}
	return true
}

// 判断一个路径是否存在,不存在则创建
func ExistsDir(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(name, os.ModeDir|0755)
			if nil != err {
				return false
			}
			return true
		}
	}
	return true
}

// 是否存在某一路径
func Fexists(ph string) bool {
	return Exists(ph)
}

// 确保某个路径的父目录存在
func FcheckParents(aph string) {
	pph := path.Dir(aph)
	err := os.MkdirAll(pph, os.ModeDir|0755)
	if nil != err {
		panic(err)
	}
}

// 读取多行属性文件，并将其变成一个 map，如果文件不存在，返回一个空 map
func Properties(ph string) map[string]string {
	pp := make(map[string]string, 20)
	var f, e_open = os.Open(ph)
	// 文件不存在，那么创建默认的的配置信息
	if f == nil || e_open != nil {
		return pp
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		var line, e_read = r.ReadString('\n')
		// 出错或者读到文件结尾，退出循环
		if e_read == io.EOF || e_read != nil {
			break
		}
		// 去掉空白
		line = strings.TrimSpace(line)

		// 注释行跳过
		if strings.HasPrefix(line, "#") {
			continue
		}
		// 开始一个多行属性处理过程，会一直读到 "#" 开始的行
		if strings.HasSuffix(line, ":") {
			key := line[:len(line)-1]
			sb := make([]string, 0, 100)
			for {
				line, e_read = r.ReadString('\n')
				if e_read == io.EOF || e_read != nil {
					break
				}
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "#") {
					break
				}
				sb = append(sb, line)
			}
			pp[key] = strings.Join(sb, "\n")
			continue
		}
		// 处理单行属性，= 作为分隔，如果没有，认为是空字符串
		ss := strings.SplitN(line, "=", 2)
		if len(ss) == 2 {
			pp[ss[0]] = ss[1]
		} else {
			pp[line] = ""
		}
	}
	return pp
}

// 文件夹中文件个数,单层目录
func DirFileNum(ph string) int {
	// 计数器
	var i = 0
	// 遍历目录,获得主动事件
	filepath.Walk(ph, func(ph string, f os.FileInfo, err error) error {
		// 文件不存在
		if f == nil {
			return nil
		}
		f.Mode()
		// 跳过文件夹
		if f.IsDir() {
			return nil
		}
		// 文件
		i++
		// 返回空
		return nil
	})
	// 返回文件总量
	return i
}

// 删除一个文件，或者文件夹，如果该路径不存在，返回 false。
// 如果是文件夹，递归删除
func RemoveAll(ph string) error {
	return os.RemoveAll(ph)
}

// 移动文件或文件夹
func Fmove(frompath, topath string) error {
	// 确保父目录存在
	FcheckParents(topath)
	// 移动
	e := os.Rename(frompath, topath)
	// 返回
	return e
}

// 通过文件结尾读取类型
func FileType(name string) string {
	// 从文件名中读取文件格式
	fileName := strings.Split(name, ".")
	// 返回
	return fileName[len(fileName)-1]
}

// 遍历目录,尝试寻找指定开头文件
func FindDirHeadFile(ph string, head string) string {
	// 保存找到的文件名称
	var name string
	// 遍历目录,获得主动事件
	filepath.Walk(ph, func(ph string, f os.FileInfo, err error) error {
		// 文件不存在
		if f == nil {
			return nil
		}
		// 跳过文件夹
		if f.IsDir() {
			return nil
		}
		// 判断文件名称是否为指定开头
		if strings.HasPrefix(f.Name(), head) {
			// 保存名称
			name = f.Name()
		}
		// 返回空
		return nil
	})
	// 返回
	return name
}

// 拷贝文件
func CopyFile(src, dst string) (err error) {
	FcheckParents(dst)
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

// 目录大小
func DirSize(path string) int64 {
	var size int64
	filepath.Walk(path, func(aph string, fi os.FileInfo, err error) error {
		if fi == nil {
			return nil
		}
		size += fi.Size()
		return nil
	})
	return size
}

// 获取路径目录
func DirName(path string) string {
	dirs := strings.Split(path, string(os.PathSeparator))
	var dir string
	for i := 0; i < len(dirs)-1; i++ {
		if i == len(dirs)-2 {
			dir += dirs[i]
		} else {
			dir += dirs[i] + string(os.PathSeparator)
		}
	}
	return dir
}
