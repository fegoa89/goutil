// Package array is a collection of functions to manipulate arrays
package array

// Inserts one or more elements to the end of the slice
func ArrayPush(slice *[]interface{}, elementsToAdd ...interface{}) int {
	*slice = append(*slice, elementsToAdd...)
	return len(*slice)
}
