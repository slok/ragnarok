/*
Package grpc will have all the proto and the generated go files of the required services
by the nodes and the masters to implement the servers and the clients
*/
package grpc // import "github.com/slok/ragnarok/grpc"

// Node status protos
//go:generate protoc -I. -I${GOOGLEPROTO_PATH} nodestatus/nodestatus.proto --go_out=plugins=grpc:.

// Failure protos
//go:generate protoc -I. failurestatus/failurestatus.proto --go_out=plugins=grpc:.
