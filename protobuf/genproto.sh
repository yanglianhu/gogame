#go${exe} install github.com/gogo/protobuf/protoc-gen-gofast@latest
./protoc --gofast_out=cs --proto_path=csdef csdef/*.proto
./protoc --gofast_out=ss --proto_path=ssdef ssdef/*.proto