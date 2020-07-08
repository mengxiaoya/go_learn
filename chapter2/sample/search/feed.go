package search

import (
	"encoding/json" // json 包提供编解码JSON 的功能
	"os"            // os 包提供访问操作系统的功能，如读文件。
)

// Go 编译器可以根据赋值运算符右边的值来推导类型，声明常量的时候不需要指定类型
// 此外，这个常量的名称使用小写字母开头，表示它只能在search包内的代码里直接访问，而不暴露到包外面。
const dataFile = "data/data.json"

// Feed 包含我们需要处理的数据源的信息
type Feed struct {
	// 每个字段的声明最后` 引号里的部分被称作标记（tag），每个标记将结构类型里字段对应到JSON 文档里指定名字的字段
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// 此函数的返回值是一组Feed类型的切片，切片是一种实现了一个动态数组的引用类型
// 第二个返回值error用来表示函数是否调用成功
// 会经常看到返回error 类型值来表示函数是否调用成功。这种用法在标准库里也很常见。
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		// 如果打开文件真的有问题，就把这个错误值返回给调用者
		return nil, err
	}

	// 当函数返回时，关闭文件
	// 关键字 defer 会安排随后的函数调用在函数返回时才执行。使用关键字defer 来安排调用Close 方法，可以保证这个函数一定会被调用。
	// 哪怕函数意外崩溃终止，也能保证关键字defer 安排调用的函数会被执行。
	// 关键字 defer 可以缩短打开文件和关闭文件之间间隔的代码行数，有助提高代码可读性，减少错误。
	defer file.Close()

	// 将文件解码到一个切片里
	// 这个切片的每一项是一个指向一个Feed 类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	// Decode 方法接受一个类型为interface{}的值作为参数。这个类型在Go 语言里很特殊，一般会配合reflect包里提供的反射功能一起使用。
	// 在这个例子里，不需要对Decode 调用之后的错误做检查。函数执行结束，这个函数的调用者可以检查这个错误值，并决定后续如何处理。

	// 这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}
