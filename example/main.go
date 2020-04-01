package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func loopHardProcess(name string) {
	defer wg.Done()
	for i := 0; i < 11; i++ {
		// consulta no banco de dados
		time.Sleep(time.Second)
		fmt.Printf("Task loop hard process %v \n", i)
	}
	fmt.Printf("Exec finish... %v \n", name)
}

func main() {
	for i := 0; i < 11; i++ {
		wg.Add(1)
		go loopHardProcess(strconv.Itoa(i))
	}
	wg.Wait()
	fmt.Println("finish main...")
}
