package gcfbd

import (
	"fmt"
	"testing"
)

func TestGCHandlerFunc(t *testing.T) {
	data := GCHandlerFunc("string", "gisaw", "petasaw")

	fmt.Printf("%+v", data)
}