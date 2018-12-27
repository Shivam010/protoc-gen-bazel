package main

import (
	"fmt"
	"github.com/lyft/protoc-gen-star"
	"time"
)

type postProcessor struct {
	author  string
	message string
}

// NewPostProcessor returns a new PostProcessor that
// adds additional information at the top of all generated files
func NewPostProcessor(author string, message ...string) pgs.PostProcessor {
	p := &postProcessor{
		author:  author,
		message: "",
	}
	for _, message := range message {
		p.message += message + "\n"
	}
	return p
}

func (p *postProcessor) Match(a pgs.Artifact) bool { return true }

func (p *postProcessor) Process(in []byte) ([]byte, error) {
	comment := fmt.Sprintf("%s# Copyright © %d %s. All rights reserved\n", p.message, time.Now().Year(), p.author)
	return append([]byte(comment), in...), nil
}
