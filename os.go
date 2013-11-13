package z

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// 获取本地MAC地址，只限Linux系统
func GetMac() string {
	var mac string
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("/sbin/ifconfig", "-a")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()
	sOut := stdout.String()
	sErr := stderr.String()
	if len(sErr) == 0 {
		rx, _ := regexp.Compile("[0-9A-F]{2}:[0-9A-F]{2}:[0-9A-F]{2}:[0-9A-F]{2}:[0-9A-F]{2}:[0-9A-F]{2}")
		macStr := rx.FindString(strings.ToUpper(sOut))
		str := strings.ToUpper(macStr)
		mac = strings.Replace(str, ":", "", -1)
	} else {
		log.Panic(sErr)
	}
	return Trim(mac)
}

// 计算一个文件的 MD5 指纹, 文件路径为磁盘绝对路径
func MD5(ph string) string {
	return Finger(md5.New(), ph)
}

// 将磁盘某个文件按照某种算法计算成加密指纹
func Finger(h hash.Hash, ph string) string {
	// 打开文件
	f, err := os.Open(ph)
	if err != nil {
		return ""
	}
	defer f.Close()
	// 读取
	io.Copy(h, bufio.NewReader(f))
	// 返回计算结果
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 对字符串进行SHA1哈希
func StrSHA1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// 通过唯一时间的字符串，返回唯一的SHA1哈希
func RandomSHA1() string {
	return StrSHA1(UnixNano())
}

// 生成一个 UUID 字符串（小写，去掉减号），需要系统支持 "uuidgen" 命令
// 返回的字符串格式如 "1694108edc6348b08364e604dee1bf35"
func UU() string {
	return strings.Replace(UU16(), "-", "", -1)
}

// 生成一个 UUID 字符串（小写），需要系统支持 "uuidgen" 命令
// 返回的字符串格式如 "1694108e-dc63-48b0-8364-e604dee1bf35"
func UU16() string {
	bs, err := exec.Command("uuidgen").Output()
	if nil != err {
		log.Fatal("fail to found command 'uuidgen' in $PATH")
	}
	return strings.ToLower(TrimBytes(bs))
}

// 解压Tar文件
func Untar(file, path string) error {
	// 打开文件
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	// 读取GZIP
	gr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gr.Close()
	// 读取TAR
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		// 打开文件
		fw := FileW(path + string(os.PathSeparator) + hdr.Name)
		// 保证文件正常关闭
		defer fw.Close()
		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
	}
	return nil
}

// 运行命令脚本，只限Linux系统
func LinuxCmd(sh string) error {
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd := exec.Command("/bin/sh", sh)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("[%s] [%s]", sh, err)
	}
	sOut := stdout.String()
	if len(sOut) != 0 {
		log.Println(sOut)
	}
	sErr := stderr.String()
	if len(sErr) != 0 {
		return fmt.Errorf(sh, sErr)
	}
	return nil
}

// 运行系统命令，只限Linux系统
func LinuxBash(sh string) error {
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd := exec.Command(sh)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("[%s] [%s]", sh, err)
	}
	sOut := stdout.String()
	if len(sOut) != 0 {
		log.Println(sOut)
	}
	sErr := stderr.String()
	if len(sErr) != 0 {
		return fmt.Errorf(sErr)
	}
	return nil
}
