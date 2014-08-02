package main

import (
	"fmt"
	"strconv"

	"github.com/robertkrimen/otto"
)

func main() {
	v := runDemo1()
	fmt.Printf("\n%v\n", v)

	v = runDemo2()
	fmt.Printf("\n%v\n", v)

	v = runDemo3(100)
	fmt.Printf("\n%v\n", v)
}

func runDemo1() otto.Value {
	return runDemo(nil, []string{"assets/demo1.js"}, "React.renderComponentToString(CommentBox({}))")
}

func runDemo2() otto.Value {
	return runDemo(nil, []string{"assets/demo2.js"}, "React.renderComponentToString(HelloWorld({}));")
}

func runDemo3(rows int) otto.Value {
	return runDemo(nil, []string{"assets/demo3.js"}, `
		var data = [];
		for (i = 0; i < `+strconv.Itoa(rows)+`; i++) {
			data.push({"id": i, "author": "Anonymous", "text": "This is comment #" + i});
		}
		React.renderComponentToString(CommentBox({data : data}));
	`)
}

func runDemo(r *renderer, files []string, cmd string) otto.Value {
	if r == nil {
		r = newRenderer()
	}

	for _, file := range files {
		r.runFile(file)
	}

	v, err := r.Run(cmd)
	if err != nil {
		panic(err)
	}
	return v
}
