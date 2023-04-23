package go_result

import (
	"testing"
)

func TestNewResult(t *testing.T) {
	aList := NewEmpty[string, any]()

	array := aList.Raw()

	if len(array) != 0 {
		t.Error("list.Raw() returned unexpected result in TestRaw")
	}
}

func TestRawPointer(t *testing.T) {
	list := NewEmpty[string, any]()

	arrayPtr := list.RawPointer()

	if arrayPtr != list.Array {
		t.Error("list.RawPointer() returned unexpected result in TestRawPointer")
	}
}

