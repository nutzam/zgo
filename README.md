zgo
===

提供一组 golang 的帮助函数集合，本项目是个代码库


# 安装

**自动安装**

	go get github.com/nutzam/zgo
	
**手动安装**

自己手动从 github 下载代码后，放置在你的 $GOPATH 的 src/github.com/nutzam/zgo 目录下

	go install github.com/nutzam/zgo
	
**安装成功的标志**

请检查你的 $GOPATH 是不是

	$GOPATH
		[src]
			[github.com]
				[nutzam]
					[zgo]           # <- 这里是下载下来的源码
						REAME.md
						disk.go
						…
		[pkg]
			[github.com]
				[nutzam]
					zgo.a           # <- 这里是编译好的包
					
# 在你的项目里使用

在你的项目里你就能正常使用这个代码库了，具体如何使用，请参看 godoc

	package main

	import (
		"fmt"
		z "github.com/nutzam/zgo"
	)
	
	func main() {
		fmt.Println(z.Ph("~/my.txt"))
	}
