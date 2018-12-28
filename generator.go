package main

import (
	"fmt"
	"github.com/lyft/protoc-gen-star"
)

// BazelBuildModule creates the bazel build file for all the proto files using protoc
type BazelBuildModule struct {
	*pgs.ModuleBase
}

// NewBazelBuildGenerator configures the BazelBuildModule with an instance of ModuleBase
func NewBazelBuildGenerator() *BazelBuildModule {
	return &BazelBuildModule{&pgs.ModuleBase{}}
}

// Name returns the module identifier
func (bm *BazelBuildModule) Name() string {
	return "BazelBuildFileGenerator"
}

// Execute uses target files and loaded packages to generate the bazel build code
func (bm *BazelBuildModule) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {

	// output will have the generated bazel build code for each target file packages
	output := make(map[string]string)

	// iterating through the targeted proto files
	for _, f := range targets {
		name, srcs, deps, tsdeps := extractDetails(f)
		dn := parentDirectory(f)

		output[dn] += createProtoLibrary(name, srcs, deps)
		output[dn] += createTsProtoLibrary(name, tsdeps)
	}

	for dn, out := range output {
		// adding package information
		out = packageInformation() + out

		// generate output file
		fn := fmt.Sprintf("%s/BUILD", dn)
		bm.AddGeneratorFile(fn, out)
	}

	return bm.Artifacts()
}

// packageInformation returns default package information to be included in the build file
// like package visibility, loading information etc
func packageInformation() string {

	// adding package default_visibility information
	out := "# The package rules can be used by any other package\n"
	out += "package(default_visibility = [\"//visibility:public\"])\n\n"

	// adding load extensions information
	out += "# Loading useful Bazel extensions\n"
	out += "load(\"@build_bazel_rules_typescript//:defs.bzl\", \"ts_proto_library\")\n\n"

	return out
}

// createProtoLibrary returns proto_library definition of the proto files
func createProtoLibrary(name string, _srcs, _deps []string) string {
	srcs := ""
	deps := ""
	for _, s := range _srcs {
		srcs += fmt.Sprintf("\t\t\"%s.proto\",\n", s)
	}
	for _, d := range _deps {
		deps += fmt.Sprintf("\t\t\"%s_proto\",\n", d)
	}

	out := fmt.Sprintf("proto_library(\n"+
		"\tname = \"%s_proto\",\n"+
		"\tsrcs = [\n%s\t],\n"+
		"\tdeps = [\n%s\t],\n)\n\n",
		name, srcs, deps,
	)
	return out
}

// createTsProtoLibrary returns ts_proto_library definition of the proto files
func createTsProtoLibrary(name string, _deps []string) string {
	deps := ""
	for _, d := range _deps {
		deps += fmt.Sprintf("\t\t\"%s_proto\",\n", d)
	}

	out := fmt.Sprintf("ts_proto_library(\n"+
		"\tname = \"%s_ts_proto\",\n"+
		"\tdeps = [\n%s\t],\n)\n\n",
		name, deps,
	)
	return out
}

// createCcProtoLibrary returns cc_proto_library definition of the proto files
func createCcProtoLibrary(name string, _deps []string) string {
	deps := ""
	for _, d := range _deps {
		deps += fmt.Sprintf("\t\t\"%s_proto\",\n", d)
	}

	out := fmt.Sprintf("cc_proto_library(\n"+
		"\tname = \"%s_cc_proto\",\n"+
		"\tdeps = [\n%s\t],\n)\n\n",
		name, deps,
	)
	return out
}
