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

// Removes repeated values in a slice
func SliceUnique(slice *[]interface{}) (uniqueslice []interface{}) {
	for _, v := range *slice {
		if !isInterfaceInSlice(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	*slice = uniqueslice
	return
}

// Returns a slice containing all the entries from slice1 that are not present in slice2.
func SliceDiff(slice1 *[]interface{}, slice2 *[]interface{}) (diffSlice []interface{}) {
	for _, v := range *slice1 {
		if !isInterfaceInSlice(v, *slice2) {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

// Returns a slice containing all the entries from slice1 that are present in slice2.
func SliceIntersect(slice1 *[]interface{}, slice2 *[]interface{}) (intersectSlice []interface{}) {
	for _, v := range *slice1 {
		if isInterfaceInSlice(v, *slice2) {
			intersectSlice = append(intersectSlice, v)
		}
	}
	return
}

// Checks if given interface exists in interface slice
func isInterfaceInSlice(inputInterface interface{}, uniqSlice []interface{}) bool {
	for _, element := range uniqSlice {
		if element == inputInterface {
			return true
		}
	}
	return false
}
