package slice

import (
	"reflect"
	"testing"
)

func TestSlicePush(t *testing.T) {
	initialSlice := make([]interface{}, 0)
	expectedSlice := []interface{}{"fox"}

	SlicePush(&initialSlice, "fox")

	if reflect.DeepEqual(initialSlice, expectedSlice) != true {
		t.Error("SlicePush() returns 'false', expected 'true'")
	}
}

func TestSlicePop(t *testing.T) {
	initialSlice := []interface{}{"red", "fox"}
	expectedSlice := []interface{}{"red"}

	SlicePop(&initialSlice)

	if reflect.DeepEqual(initialSlice, expectedSlice) != true {
		t.Error("SlicePop() expected %v, got %v", expectedSlice, initialSlice)
	}
}

func TestSliceShift(t *testing.T) {
	initialSlice := []interface{}{"red", "fox"}
	expectedSlice := []interface{}{"fox"}

	SliceShift(&initialSlice)

	if reflect.DeepEqual(initialSlice, expectedSlice) != true {
		t.Error("SliceShift() expected %v, got %v", expectedSlice, initialSlice)
	}
}

func TestRemoveItemByIndex(t *testing.T) {
	initialSlice := []interface{}{"red", "fox"}
	expectedSlice := []interface{}{"fox"}

	RemoveItemByIndex(&initialSlice, 0)

	if reflect.DeepEqual(initialSlice, expectedSlice) != true {
		t.Error("RemoveItemByIndex() expected %v, got %v", expectedSlice, initialSlice)
	}
}
