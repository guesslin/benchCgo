package mybench

/*
#cgo CFLAGS: -O3

#include "m.h"
*/
import "C"

import (
	"fmt"
	"time"
)

func MyTest(gotimes int) {
	fmt.Println("no C to GO callback bench test start")
	times := C.int(gotimes)
	var sum C.int
	start := time.Now()
	sum = C.non_callback_bench(times)
	end := time.Now()
	fmt.Println("Duration:", end.Sub(start), "sum", int(sum), "times", int(times))
	fmt.Println("no C to GO callback bench test end")
	sum = C.int(0)
}

func MyTestCallback(gotimes int) {
	times := C.int(gotimes)
	var sum C.int
	fmt.Println("C to GO callback bench test start")
	start := time.Now()
	sum = C.callback_bench(times)
	end := time.Now()
	fmt.Println("Duration:", end.Sub(start), "sum", int(sum), "times", int(times))
	fmt.Println("C to GO callback bench test end")
	sum = C.int(0)
}

func MyTestPure(gotimes int) {
	times := C.int(gotimes)
	var sum C.int
	fmt.Println("pure GO bench test start")
	start := time.Now()
	sum = bench(times)
	end := time.Now()
	fmt.Println("Duration:", end.Sub(start), "sum", int(sum), "times", int(times))
	fmt.Println("pure GO bench test end")
}

//export my_go_sum
func my_go_sum(x, y C.int) C.int {
	return x + y
}

func bench(times C.int) C.int {
	var sum C.int = 0
	var i C.int
	for i = 0; i < times; i++ {
		sum += my_go_sum(i, i+1)
	}
	return sum
}
