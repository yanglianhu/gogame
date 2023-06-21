module gatewaysvr

go 1.18

replace gogame/protobuf/cs => ../../protobuf/cs

require (
	gogame/protobuf/cs v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)
