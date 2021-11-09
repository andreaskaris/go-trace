package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; ; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
}
