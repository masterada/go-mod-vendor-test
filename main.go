package main

import (
	"k8s.io/apimachinery/pkg/runtime/serializer/protobuf"
)

func main() {
	protobuf.IsNotMarshalable(nil)
}
