package slice

import (
	"reflect"
	"testing"
)

func TestSlicePush(t *testing.T) {
	initialSlice := []interface{}{}
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
		t.Errorf("SlicePop() expected %v, got %v", expectedSlice, initialSlice)
	}
}

func TestSliceShift(t *testing.T) {
	initialSlice := []interface{}{"red", "fox"}
	expectedSlice := []interface{}{"fox"}

	SliceShift(&initialSlice)

	if reflect.DeepEqual(initialSlice, expectedSlice) != true {
		t.Errorf("SliceShift() expected %v, got %v", expectedSlice, initialSlice)
	}
}

func TestRemoveItemByIndex(t *testing.T) {
	initialSlice := []interface{}{"red", "fox"}
	expectedSlice := []interface{}{"fox"}

	RemoveItemByIndex(&initialSlice, 0)

	if reflect.DeepEqual(initialSlice, expectedSlice) != true {
		t.Errorf("RemoveItemByIndex() expected %v, got %v", expectedSlice, initialSlice)
	}
}

func TestSliceUnique(t *testing.T) {
	initialSlice := []interface{}{"red", "red", "red", "fox"}
	expectedSlice := []interface{}{"red", "fox"}

	SliceUnique(&initialSlice)

	if reflect.DeepEqual(initialSlice, expectedSlice) != true {
		t.Errorf("SliceUnique() expected %v, got %v", expectedSlice, initialSlice)
	}
}

func TestSliceDiff(t *testing.T) {
	firstSlice := []interface{}{"red", "dog", "fox"}
	secondSlice := []interface{}{"red", "fox"}
	expectedSlice := []interface{}{"dog"}

	diffSlice := SliceDiff(&firstSlice, &secondSlice)

	if reflect.DeepEqual(diffSlice, expectedSlice) != true {
		t.Errorf("SliceDiff() expected %v, got %v", expectedSlice, diffSlice)
	}
}

func TestSliceIntersect(t *testing.T) {
	firstSlice := []interface{}{"red", "dog", "fox"}
	secondSlice := []interface{}{"dog"}
	expectedSlice := []interface{}{"dog"}

	diffSlice := SliceIntersect(&firstSlice, &secondSlice)

	if reflect.DeepEqual(diffSlice, expectedSlice) != true {
		t.Errorf("SliceIntersect() expected %v, got %v", expectedSlice, diffSlice)
	}
}
