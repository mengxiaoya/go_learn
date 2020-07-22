package main

import (
	"fmt"
	"time"

	drum "github.com/openex27/drumstick"
)

func sumEcho() {
	fmt.Printf("time:%s, task...\n", time.Now())
}

func main() {
	task, err := drum.NewTask(time.Now(), 1*time.Second, sumEcho)
	if err != nil {
		panic(err)
	}
	task.Start()
	time.Sleep(10 * time.Second)
	task.Reset(time.Now(), 2*time.Second)
	time.Sleep(10 * time.Second)
	task.Stop()
}
