package main

import (
	"fmt"
	"github.com/lyft/protoc-gen-star"
)

type postProcessor string

// NewPostProcessor returns a new PostProcessor that
// adds additional information at the top of all generated files
func NewPostProcessor(message string) pgs.PostProcessor {
	p := postProcessor(message)
	return &p
}

func (p *postProcessor) Match(a pgs.Artifact) bool { return true }

func (p *postProcessor) Process(in []byte) ([]byte, error) {
	comment := fmt.Sprintf("# %s\n\n", *p)
	return append([]byte(comment), in...), nil
}
