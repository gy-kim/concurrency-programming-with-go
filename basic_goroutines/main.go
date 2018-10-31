package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	godur, _ := time.ParseDuration("10ms")
	runtime.GOMAXPROCS(2)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
