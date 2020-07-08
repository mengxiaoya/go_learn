package search

import (
	"log"
	"sync" // sync包提供同步goroutine的功能
)

// 声明一个map，key为string类型，value为自定义的Matcher类型
var matchers = make(map[string]Matcher)

// run方法执行搜索逻辑
// 函数定义：func 方法名(参数列表) (返回值)
func Run(searchTerm string) {
	// 获取需要搜索的数据源列表，其实就是读文件
	// := 是简化变量声明运算符，简化变量声明运算符只是一种简化记法，让代码可读性更高。它使用关键字 var 声明的变量没有任何区别
	// 所以如下紧挨的两行代码是等价的

	// 根据经验，如果需要声明初始值为零值的变量，应该使用 var 关键字声明变量；
	// 如果提供确切的非零值初始化变量或者使用函数返回值创建变量，应该使用简化变量声明运算符
	feeds, err := RetrieveFeeds()
	// var feeds, err = RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果
	// go语言中通道和映射map和切片slice一样，也是引用类型，不过
	// 通道本身实现的是一组带类型的值，这组值用于在 goroutine 之间传递数据。通道内置同步机制，从而保证通信安全。
	results := make(chan *Result)

	// 下面两行代码是为了防止程序在全部搜索执行完之前终止
	// 构造一个 waitGroup，以便处理所有的数据源
	// 这个程序使用 sync 包的 WaitGroup 跟踪所有启动的 goroutine。非常推荐使用 WaitGroup 来
	// 跟踪 goroutine 的工作是否完成。WaitGroup 是一个计数信号量，我们可以利用它来统计所有的
	// goroutine 是不是都完成了工作。
	var waitGroup sync.WaitGroup
	// 设置需要等待处理每个数据源的 goroutine 的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	// 关键字 range 可以用于迭代数组、字符串、切片、映射和通道。
	// 使用 for range 迭代切片时，每次迭代会返回两个值。第一个值是迭代的元素在切片里的索引位置，第二个值是元素值的一个副本。

	// 这里的下划线标识符的作用是占位符，占据了保存 range 调用返回的索引值的变量的位置。如果
	// 要调用的函数返回多个值，而又不需要其中的某个值，就可以使用下划线标识符将其忽略。如果需要的话
	// 则定义具体的变量来接收。
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		// 查找 map 里的键时，有两个选择：要么赋值给一个变量，要么为了精确查找，赋值给两个变量。
		// 赋值给两个变量时第一个值和赋值给一个变量时的值一样，是 map 查找的结果值。
		// 如果指定了第二个值，就会返回一个布尔标志，来表示查找的键是否存在于 map 里。
		// 如果这个键不存在，map 会返回其值类型的零值作为返回值，如果这个键存在，map 会返回键所对应值的副本。
		matcher, exists := matchers[feed.Type]
		if !exists {
			// 如果不存在，使用默认匹配器。这样程序在不知道对应数据源的具体类型时，也可以执行，而不会中断。
			matcher = matchers["default"]
		}

		//var matcher = matchers[feed.Type]
		//if matcher == nil {
		//	matcher = matchers["default"]
		//}

		// 启动一个 goroutine 来执行搜索，这样可以并发地独立处理每个数据源的数据。
		// 使用关键字 go 启动一个 goroutine，并对这个 goroutine 做并发调度。
		// 这里使用关键字 go 启动了一个匿名函数作为 goroutine
		// 变量 feed 是一个指 针变量。指针变量可以方便地在函数之间共享数据。使用指针变量可以让函数访问并修改一个变
		// 量的状态，而这个变量可以在其他函数甚至是其他 goroutine 的作用域里声明。

		// 在 Go 语言中，所有的变量都以值的方式传递。因为指针变量的值是所指向的内存地址，在函数间传递指针变量，是在传递
		// 这个地址值，所以依旧被看作以值的方式在传递。
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done() // 一旦 Match 函数调用完毕，递减 WaitGroup 的计数。
			// 这里还有一个点就是waitGroup没有作为参数传入到匿名函数中，但是还能访问到
			// 这是因为Go语言支持闭包，函数可以直接访问到那些没有作为参数传入
			// 的变量。匿名函数并没有拿到这些变量的副本，而是直接访问外层函数作用域中声明的这些变量
			// 本身。因为 matcher 和 feed 变量每次调用时值不相同，所以并没有使用闭包的方式访问这两个变量。
		}(matcher, feed)
	}

	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func() {
		// 等候所有任务完成，这个方法会造成阻塞，直到WaitGroup内部的计数到达0
		waitGroup.Wait()

		//  用关闭通道的方式，通知 Display 函数，可以退出程序了
		close(results)
	}()

	// 该函数的作用就是打印返回的结果，并且在最后一个结果显示完后返回
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
