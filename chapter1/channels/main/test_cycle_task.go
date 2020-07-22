package main

import (
	"fmt"
	"time"
)

func init() {
	//程序初始化的时候就初始化调用一次
	do()
}

func main() {
	//主程序不能挂
	time.Sleep(10 * time.Second)
}

var count = 2

func do() {
	//t := time.Now()
	////获取下一次执行的时间，可以在年月日时分秒上加，这里设置1分钟执行一次
	//next := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()+1, 0, 0, SysTimeLocation)
	//fmt.Printf("next  type: %T,\t val: %v\n", next, next)
	////获取下次执行时间与当前时间的差
	//duration := next.Sub(time.Now())
	//fmt.Printf("duration  type: %T,\t val: %v\n", duration, duration)

	/*预约下次执行执行计划，因为在程序初始化的时候已经调用了do()方法，
	*在do()每次执行完，都会再预约下次执行计划，直到主程序die*/
	count++
	time.AfterFunc(time.Duration(1)*time.Second, do)
	fmt.Println("****")
}
