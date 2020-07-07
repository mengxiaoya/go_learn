package main

import (
	// log和os为go的标准库的包，从go标准库导入包时只需要导入包名即可，编译器查找包时总是会到GOROOT和GOPATH对应的位置去找
	// 第三方包的导入例如下面的matchers和search，还需要指定路径
	"log" // log包提供打印日志信息到标准输出、标准错误或者自定义设备
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

// 在 Go 语言中，如果 main 函数返回，整个程序也就终止了。Go 程序终止时，还会关闭所有
// 之前启动且还在运行的 goroutine。写并发程序的时候，最佳做法是，在 main 函数返回前，清理
// 并终止所有之前启动的 goroutine。编写启动和终止时的状态都很清晰的程序，有助减少 bug，防
// 止资源异常。
func main() {
	// 当代码导入了一个包时，程序可以直接访问这个包中任意一个公开的标识符，这些标识符以大写字母开头，以小写字母开头
	// 的标识符是不能被其他包中的代码直接访问。但是可以间接访问，例如一个函数返回一个未公开类型的值，那么这个函数的
	// 任何调用者，哪怕调用者不是在这个包里声明的，都可以访问这个值。
	search.Run("president")
}
