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
BenchmarkRender1	      10	 132139860 ns/op
BenchmarkRender5	      10	 154325580 ns/op
BenchmarkRender10	      10	 184817942 ns/op
BenchmarkRender20	       5	 265151415 ns/op
BenchmarkRender50	       2	 547846865 ns/op
BenchmarkRender100	     1	1241521308 ns/op
BenchmarkRender200	     1	3666366314 ns/op
BenchmarkRender500	     1 18904415122 ns/op
update
ok  	github.com/101loops/go-reactjs	32.312s
```