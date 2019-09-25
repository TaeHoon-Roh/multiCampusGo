package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func mytest() func(myTemp int) int {
	a := make([]int, 100000)
	b := make([]int, 100000)

	a = b
	b = a
	// return a + b
	return func(myTemp int) int {
		// return ()
		a = append(a, b...)

		return 0
	}

}

func mytestCopy() int {
	a := make([]int, 10000)
	b := make([]int, 10000)

	a = b
	b = a
	return 0
}

func main() {
	var wait sync.WaitGroup
	wait.Add(1)
	go PrintMemUsage()

	myFunc := mytest()
	fmt.Println(myFunc)
	funResult := myFunc(1)
	for i := 0; i < 10; i++ {
		myFunc(i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println(funResult)

	wait.Wait()

}

func PrintMemUsage() {
	var m runtime.MemStats
	for {
		runtime.ReadMemStats(&m)
		// For info on each, see: https://golang.org/pkg/runtime/#MemStats
		fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
		fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
		fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
		fmt.Printf("\tNumGC = %v\n", m.NumGC)

		time.Sleep(time.Second * 1)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
