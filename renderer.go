package main

import (
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

type renderer struct {
	*otto.Otto
}

func newRenderer() *renderer {
	r := &renderer{otto.New()}
	r.runFile("assets/global.js")
	r.runFile("assets/react.js")
	return r
}

func (r *renderer) runFile(path string) otto.Value {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	result, err := r.Run(data)
	if err != nil {
		panic(err)
	}

	return result
}
