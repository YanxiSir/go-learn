package main

import (
	"fmt"
	"sync"
	"time"
)

//WaitGroup test
func main() {
	var n sync.WaitGroup
	for i := 0; i < 20; i++ {
		n.Add(1)
		go func(i int, n *sync.WaitGroup) {
			defer n.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("goroutine %d is running \n", i)
		}(i, &n)
	}
	n.Wait()
	fmt.Println("finished")
}
