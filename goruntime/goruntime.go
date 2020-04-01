package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup


func hardTask(name string) {
	defer wg.Done()
	for i := 0; i < 11; i++ {
		time.Sleep(1 + time.Second)
		fmt.Printf("Hard task %s...\n", name)
	}
	fmt.Printf(" Hard task %s DONE \n", name)
}

func easyTask() {
	time.Sleep(1 + time.Second)
	fmt.Println("========== Easy Task ==========")
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hardTask(strconv.Itoa(i))
	}

	easyTask()
	wg.Wait()
	fmt.Println("Done Finish!!!")
}
