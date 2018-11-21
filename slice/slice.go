// Package slice is a collection of functions to manipulate slices
package slice

// Inserts one or more elements to the end of the slice
func SlicePush(slice *[]interface{}, elementsToAdd ...interface{}) int {
	*slice = append(*slice, elementsToAdd...)
	return len(*slice)
}

// Removes the last element from an slice and returns that removed element
func SlicePop(slice *[]interface{}) interface{} {
	if len(*slice) == 0 {
		return nil
	}

	*slice = (*slice)[:len(*slice)-1]
	return (*slice)[len(*slice)-1]
}

// Removes the first element from an slice and returns that removed element
func SliceShift(slice *[]interface{}) interface{} {
	if len(*slice) == 0 {
		return nil
	}

	*slice = (*slice)[1:]
	return (*slice)[0]
}

// Remove an item by index position
func RemoveItemByIndex(slice *[]interface{}, position int) int {
	*slice = append((*slice)[:position], (*slice)[position+1:]...)
	return len(*slice)
}
