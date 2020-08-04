package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

type TestJob struct {
	startTime string
	endTime   string
}

func (this TestJob) Run() {
	fmt.Println("testJob1...", this.startTime, this.endTime)
}

func main() {
	//go func() {
	testCron()
	//}()
	log.Printf("after test cron")

}

func hello(time string) {
	log.Println("cron running!", time)
}

func testCron() {
	c := cron.New(cron.WithSeconds())
	spec := "* * * * * * 2020"
	//c.AddFunc(spec, func() {
	//	log.Println("cron running!")
	//})
	c.AddJob(spec, TestJob{"2018", "2019"}) // 可传参的
	//entryID, _ := c.AddJob(spec, TestJob{"2018", "2019"}) // 可传参的
	//
	go func() {
		c.Start()
	}()

	//select {
	//case <-time.Tick(3 * time.Second):
	//	log.Println("190hihjkgaiug")
	//	log.Printf("3 entryID: %d, ", entryID)
	//	c.Start()
	//}

	c.AddJob(spec, TestJob{"111", "222"})
	time.Sleep(1000000000)
	//defer c.Stop()

	//
	//select {
	//case <-time.Tick(10 * time.Second):
	//	log.Println("0720dingshiqi1")
	//	log.Printf("5 entryID: %d, ", entryID)
	//	c.Stop()
	//}
}
