package search

// match.go 代码文件包含创建不同类型匹配器的代码

import (
	"log"
)

// 声明一个结构类型
type Result struct {
	Field   string
	Content string
}

// 声明一个接口类型
// 命名接口的时候，也需要遵守Go 语言的命名惯例。如果接口类型只包含一个方法，那么这个接口类型的名字以er 结尾。
type Matcher interface {
	// 只声明了一个Search 方法，这个方法输入一个指向Feed 类型值的指针和一个string 类型的搜索项
	// 这个方法返回两个值：一个指向Result 类型值的指针的切片，另一个是错误值
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed to run
// searches concurrently.
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// Perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}
}

// Display writes results to the console window as they
// are received by the individual goroutines.
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
