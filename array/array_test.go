package array

import (
	"reflect"
	"testing"
)

func TestArrayPush(t *testing.T) {
	initialSlice := make([]interface{}, 0)
	expectedSlice := []interface{}{"fox"}

	ArrayPush(&initialSlice, "fox")

	if reflect.DeepEqual(initialSlice, expectedSlice) != true {
		t.Error("TestArrayPush() returns 'false', expected 'true'")
	}
}
