React.js server-side rendering with Go
=====

This experiment is based on [otto](https://github.com/robertkrimen/otto), a Javascript interpreter in Go.

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
BenchmarkRender1	      10	  133332794 ns/op
BenchmarkRender5	      10	  155127450 ns/op
BenchmarkRender10	      10	  195377940 ns/op
BenchmarkRender20	       5	  285274908 ns/op
BenchmarkRender50	       2	  591757756 ns/op
BenchmarkRender100	     1	 1844587412 ns/op
BenchmarkRender200	     1	 4058467653 ns/op
BenchmarkRender500	     1	18851352352 ns/op

ok  	github.com/101loops/go-reactjs	33.749s
```