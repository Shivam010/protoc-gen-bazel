# protoc-gen-bazel
protoc-gen-bazel is a Bazel build file generator which is used to convert protocol buffers files to corresponding typescipts files.

Installation
---
``` 
go get -u github.com/Shivam010/protoc-gen-bazel/...
go install github.com/Shivam010/protoc-gen-bazel
```

Command
---
Command to generate the build files for the examples-protos:
```
protoc -I . examples/person/person.proto examples/address/address.proto examples/address/zip_code.proto --bazel_out .
```
This command will generate Bazel build files in the same location as that of the corresponding proto.

Now, update the ```WORKSPACE``` file if required and then, start the Bazel build, to generate the typescript files.

```
bazel build //...:*
```
