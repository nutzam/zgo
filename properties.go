package z

import (
	"io"
	"os"
	"strings"
)

// 根据一个文件路径，解析 Nutz 的多行 properties 文件。
// 这个文件必须是 utf-8
func PPreadf(ph string) map[string]string {
	f, _ := os.Open(Ph(ph))
	return PPread(f)
}

// 从一个流中解析一个 Nutz 的多行 properties 文件内容
// 这个流必须是 utf-8
func PPread(r io.Reader) (pp map[string]string) {
	// 准备返回值
	pp = make(map[string]string)
	// 读取文件全部内容
	str, _ := Utf8r(r)

	// fmt.Println(str)
	// fmt.Println(strings.Repeat("=", 80))

	// 逐行分析
	lines := strings.Split(str, "\n")
	N := len(lines)
	for i := 0; i < N; i++ {
		line := lines[i]
		s := strings.TrimSpace(line)
		// 注释行
		switch {
		// 空行
		case len(s) == 0:
			continue
		// 注释行
		case strings.HasPrefix(s, "#"):
			continue
		// 多行模式
		case strings.HasSuffix(s, ":"):
			key := string(s[:len(s)-1])
			from := i
			// 持续读取
			for ; i < N; i++ {
				line = lines[i]
				if strings.HasPrefix(line, "#") {
					break
				}
			}
			val := strings.Join(lines[from:i], "\n")
			pp[key] = val
		// 默认为名值对
		default:
			flds := strings.SplitN(s, "=", 2)
			if len(flds) == 1 {
				pp[strings.TrimSpace(flds[0])] = ""
			} else {
				pp[strings.TrimSpace(flds[0])] = strings.TrimSpace(flds[1])
			}
		}
	}

	return
}
