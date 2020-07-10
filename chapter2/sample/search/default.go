package search

// defaultMatcher 实现了默认匹配器
// 我们使用一个空结构声明了一个名叫defaultMatcher 的结构类型。空结构
// 在创建实例时，不会分配任何内存。这种结构很适合创建没有任何状态的类型。对于默认匹配器
// 来说，不需要维护任何状态，所以我们只要实现对应的接口就行。
type defaultMatcher struct{}

// init 函数将默认匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 实现了默认匹配器的行为
// Search 方法的声明也声明了defaultMatcher 类型的值的接收者
// 这里声明的函数Search前面的(m defaultMatcher)就代表接受者

// 如果声明函数的时候带有接收者，则意味着声明的这个方法会和指定的接收者的
// 类型绑在一起。在我们的例子里，Search 方法与defaultMatcher 类型的值绑在一起。这意
// 味着我们可以使用defaultMatcher 类型的值或者指向这个类型值的指针来调用Search 方
// 法。无论我们是使用接收者类型的值来调用这个方，还是使用接收者类型值的指针来调用这个
// 方法，编译器都会正确地引用或者解引用对应的值，作为接收者传递给Search 方法

// 因为大部分方法在被调用后都需要维护接收者的值的状态，所以，一个最佳实践是，将方法
// 的接收者声明为指针。对于defaultMatcher 类型来说，使用值作为接收者是因为创建一个
// defaultMatcher 类型的值不需要分配内存。由于defaultMatcher 不需要维护状态，所以
// 不需要指针形式的接收者。

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

// 与直接通过值或者指针调用方法不同，如果通过接口类型的值调用方法，规则有很大不同
// 例如下面的代码段1 和代码段2

// 代码段1
// 方法声明为使用指向 defaultMatcher 类型值的指针作为接收者
// func (m *defaultMatcher) Search(feed *Feed, searchTerm string)
// 通过 interface 类型的值来调用方法
// var dm defaultMatcher
// var matcher Matcher = dm     // 将值赋值给接口类型
// matcher.Search(feed, "test") // 使用值来调用接口方法

// 代码段2
// 方法声明为使用 defaultMatcher 类型的值作为接收者
// func (m defaultMatcher) Search(feed *Feed, searchTerm string)
// 通过 interface 类型的值来调用方法
// var dm defaultMatcher
// var matcher Matcher = &dm // 将指针赋值给接口类型
// matcher.Search(feed, "test") // 使用指针来调用接口方法

// 代码段1执行go build编译的时候会报错：
// cannot use dm (type defaultMatcher) as type Matcher in assignment
// 而代码段2可以正常编译，这是因为
// 使用指针作为接收者声明的方法，只能在接口类型的值是一个指针的时候被调用。
// 使用值作为接收者声明的方法，在接口类型的值为值或者指针时，都可以被调用。
