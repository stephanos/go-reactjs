package main

import (
	"strconv"
	"testing"
)

func benchmarkRender(i int, b *testing.B) {
	r := newRenderer([]string{"assets/demo3.js"})
	for n := 0; n < b.N; n++ {
		r.runCmd(`
			var data = [];
			for (i = 0; i < ` + strconv.Itoa(i) + `; i++) {
				data.push({"id": i, "author": "Anonymous", "text": "This is comment #" + i});
			}
			React.renderComponentToString(CommentBox({data : data}));
		`)
	}
}

func BenchmarkRender1(b *testing.B)   { benchmarkRender(1, b) }
func BenchmarkRender5(b *testing.B)   { benchmarkRender(5, b) }
func BenchmarkRender10(b *testing.B)  { benchmarkRender(10, b) }
func BenchmarkRender20(b *testing.B)  { benchmarkRender(20, b) }
func BenchmarkRender50(b *testing.B)  { benchmarkRender(50, b) }
func BenchmarkRender100(b *testing.B) { benchmarkRender(100, b) }
func BenchmarkRender200(b *testing.B) { benchmarkRender(200, b) }
func BenchmarkRender500(b *testing.B) { benchmarkRender(500, b) }
