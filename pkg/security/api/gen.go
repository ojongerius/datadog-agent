//go:generate protoc -I. --go_out=plugins=grpc,paths=source_relative:. api.proto

package api
