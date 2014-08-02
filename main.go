package main

import "fmt"

func main() {
	v := newRenderer([]string{"assets/demo1.js"}).
		runCmd("React.renderComponentToString(CommentBox({}))")
	fmt.Printf("\n%v\n", v)

	v = newRenderer([]string{"assets/demo2.js"}).
		runCmd("React.renderComponentToString(HelloWorld({}))")
	fmt.Printf("\n%v\n", v)

	v = newRenderer([]string{"assets/demo3.js"}).
		runCmd(`
			var data = [
				{"id": 0, "author": "Anonymous", "text": "This is a comment"},
				{"id": 1, "author": "Anonymous", "text": "This is another comment"},
			]
			React.renderComponentToString(CommentBox({data : data}));
		`)
	fmt.Printf("\n%v\n", v)
}
