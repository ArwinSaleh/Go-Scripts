package main

import (
	"fmt"
	"time"
)

func main()  {
	var count int = 10

	for i := 0; i < count; i++ {

		fmt.Println("TESTING")
		time.Sleep(1 * time.Second)
	}
}