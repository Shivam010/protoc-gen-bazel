# Bazel build file is generated by protoc-gen-bazel. DO NOT EDIT.

# The package rules can be used by any other package
package(default_visibility = ["//visibility:public"])

# Loading useful Bazel extensions
proto_library(
	name = "address_proto",
	srcs = [
		"address.proto",
	],
	deps = [
		"//examples/address:zip_code_proto",
	],
)

ts_proto_library(
	name = "address_ts_proto",
	deps = [
		"//examples/address:address_proto",
	],
)

proto_library(
	name = "zip_code_proto",
	srcs = [
		"zip_code.proto",
	],
	deps = [
	],
)

ts_proto_library(
	name = "zip_code_ts_proto",
	deps = [
		"//examples/address:zip_code_proto",
	],
)

