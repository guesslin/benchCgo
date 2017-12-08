/*
result
$ ./benchCgo
no C to GO callback bench test start
Duration: 1.068µs sum 891896832 times 500000
no C to GO callback bench test end
pure GO bench test start
Duration: 339.836µs sum 891896832 times 500000
pure GO bench test end
C to GO callback bench test start
Duration: 98.770956ms sum 891896832 times 500000
C to GO callback bench test end
*/
package main

import (
	"github.com/guesslin/benchCgo/mybench"
	"runtime"
)

func main() {
	times := 500000
	mybench.MyTest(times)
	runtime.Gosched()
	mybench.MyTestPure(times)
	runtime.Gosched()
	mybench.MyTestCallback(times)
}
