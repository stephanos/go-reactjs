package main

import "fmt"

func main() {
	v := newRenderer([]string{"assets/demo1.js"}).
		runCmd(`
			var component = React.createElement(CommentBox, { foo: 'bar' });
			React.renderToString(component);
		`)
	fmt.Printf("\n%v\n", v)

	v = newRenderer([]string{"assets/demo2.js"}).
		runCmd(`
			var component = React.createElement(HelloWorld, { foo: 'bar' });
			React.renderToString(component);
		`)
	fmt.Printf("\n%v\n", v)

	v = newRenderer([]string{"assets/demo3.js"}).
		runCmd(`
			var data = [
				{"id": 0, "author": "Anonymous", "text": "This is a comment"},
				{"id": 1, "author": "Anonymous", "text": "This is another comment"},
			]
			var component = React.createElement(CommentBox, {data : data});
			React.renderToString(component);
		`)
	fmt.Printf("\n%v\n", v)
}
