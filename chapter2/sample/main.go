package main

import (
	"log"
	"os"

	// _ 的作用是：让Go语言对包做初始化操作，但是并不使用包里的标识符
	// 因为为了让程序的可读性更强，Go编译器不允许声明导入某个包却不使用，所以这个下划线让编译器接收这类导入
	// 并且调用对应包内的所有代码文件里定义的init函数，对应到这个的话就是调用matchers包下的rss.go中的init函数
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init函数在main函数之前调用
func init() {
	// 下面这行代码的作用是将标准库里日志类的输出从默认的标准错误stderr设置为标准输出stdout
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
