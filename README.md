# Target

How many callback lost when c code call go code?

# How to run?

	$ go build
	$ ./benchCgo

# Result

```
$ ./benchCgo
no C to GO callback bench test start
Duration: 3.899µs sum 891896832 times 500000
no C to GO callback bench test end

pure GO bench test start
Duration: 313.832µs sum 891896832 times 500000
pure GO bench test end

GO to C callback bench test start
Duration: 294.42µs sum 891896832 times 500000
GO to C callback bench test end

C to GO callback bench test start
Duration: 106.095744ms sum 891896832 times 500000
C to GO callback bench test end
```
