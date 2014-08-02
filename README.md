React.js server-side rendering with Go
=====

This experiment is based on [otto](https://github.com/robertkrimen/otto), a Javascript interpreter for Go.

To run:
```bash
go build ./...
./go-reactjs-benchmark
```

To benchmark:
```bash
go test -bench=.
```

The results on my MacBook Pro (2.4 GHz i5):
```bash
BenchmarkRender1	     100	  17128739 ns/op
BenchmarkRender5	      50	  47324904 ns/op
BenchmarkRender10	      20	  79839996 ns/op
BenchmarkRender20	      10	 164226676 ns/op
BenchmarkRender50	       2	 612836671 ns/op
BenchmarkRender100	     1	1777275883 ns/op
BenchmarkRender200	     1	4190131936 ns/op
BenchmarkRender500	     1	20789000942 ns/op

ok  	github.com/101loops/go-reactjs	39.314s
```