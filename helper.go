package main

import (
	"fmt"
	"strings"
	"github.com/lyft/protoc-gen-star"
)

func fileName(f pgs.File) string {
	fn := f.Name().Split()[0]
	fparts := strings.Split(fn, "/")

	// excluding file path
	return fparts[len(fparts)-1]
}

func parentDirectory(f pgs.File) string {
	path := f.Name().Split()[0]
	folds := strings.Split(path, "/")

	return strings.Join(folds[:len(folds)-1], "/")
}

func extractDetails(f pgs.File) (name string, srcs, deps, tsdeps []string) {
	name = fileName(f)

	// proto file is itself its source and dependency
	srcs = append(srcs, name)
	_d := fmt.Sprintf("//%s:%s", parentDirectory(f), name)
	tsdeps = append(tsdeps, _d)

	// handling dependencies (i.e. imports)
	for _, imp := range f.Imports() {

		switch imp.FullyQualifiedName() {

		case ".google.protobuf":
			d := fmt.Sprintf("@com_google_protobuf//:%s", fileName(imp))
			deps = append(deps, d)

		default:
			d := fmt.Sprintf("//%s:%s", parentDirectory(imp), fileName(imp))
			deps = append(deps, d)

		}
	}

	return
}

