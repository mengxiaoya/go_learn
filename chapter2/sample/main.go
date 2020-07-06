package main

import (
	"log"
	"os"

	// _ 的作用是：让Go语言对包做初始化操作，但是并不使用包里的标识符
	// 因为为了让程序的可读性更强，Go编译器不允许声明导入某个包却不使用
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
