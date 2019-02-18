package main

import (
	"runtime"
)

func main() {
	pool := New(10)
	println(runtime.NumGoroutine())
	for i := 0; i < 1000; i++ {
		pool.Add(1)
		go func() {
			println(runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	println(runtime.NumGoroutine())
}
