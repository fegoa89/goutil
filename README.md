# goutil
Collection of small reusable Go functions - WIP

## Install

Install with
```sh
go get github.com/fegoa89/goutils
```

## Table of Contents

### String

<details>
<summary>View contents</summary>

* [`Capitalize`](#Capitalize)
* [`CapitalizeTitle`](#CapitalizeTitle)
* [`UpperCamelCase`](#UpperCamelCase)
* [`LowerCamelCase`](#LowerCamelCase)
* [`ToSnakeCase`](#ToSnakeCase)
* [`ToKebabCase`](#ToKebabCase)
* [`Humanize`](#Humanize)
* [`Truncate`](#Truncate)
* [`IsEmptyOrBlank`](#IsEmptyOrBlank)
* [`IsBlank`](#IsBlank)
* [`IsEmpty`](#IsEmpty)
* [`DeleteWhitespace`](#DeleteWhitespace)
* [`CaseInsensitiveEquals`](#CaseInsensitiveEquals)
* [`IsBoolean`](#IsBoolean)
* [`IsInteger`](#IsInteger)
* [`IsFloat`](#IsFloat)
* [`IndexOf`](#IndexOf)
</details>

### Slice

<details>
<summary>View contents</summary>

* [`SlicePush`](#SlicePush)
* [`SlicePop`](#SlicePop)
* [`SliceShift`](#SliceShift)
* [`RemoveItemByIndex`](#RemoveItemByIndex)
* [`SliceUnique`](#SliceUnique)
* [`SliceDiff`](#SliceDiff)
* [`SliceIntersect`](#SliceIntersect)
</details>

---

## String

Import
```go
import "github.com/fegoa89/goutils/string"
```

### Capitalize

Capitalizes the first letter of a string.
```golang
Capitalize("hello") // Hello
```

<details>
<summary>Examples</summary>

```golang
Capitalize("my string") // My string
```

</details>

### CapitalizeTitle
Capitalizes string excluding words such as "and", "a", "with", "or".
You can provide an optional string slice with words that will be excluded too
```golang
CapitalizeTitle("foo and bar") // Foo and Bar
```

### UpperCamelCase
Convert string to UpperCamelCase
```golang
UpperCamelCase("foo and bar") // FooAndBar
```

<details>
<summary>Examples</summary>

```golang
UpperCamelCase("foo and bar") // FooAndBar
UpperCamelCase("foo.and.bar") // FooAndBar
UpperCamelCase("foo-and-bar") // FooAndBar
UpperCamelCase("foo_and_bar") // FooAndBar
UpperCamelCase("Foo and bar") // FooAndBar
```

</details>

### LowerCamelCase
Convert string to LowerCamelCase
```golang
LowerCamelCase("foo and bar") // fooAndBar
```

<details>
<summary>Examples</summary>

```golang
LowerCamelCase("foo and bar") // fooAndBar
LowerCamelCase("foo.and.bar") // fooAndBar
LowerCamelCase("foo-and-bar") // fooAndBar
LowerCamelCase("foo_and_bar") // fooAndBar
LowerCamelCase("Foo and bar") // fooAndBar
```

</details>

### ToSnakeCase
Convert string to to SnakeCase
```golang
ToSnakeCase("fooAndBar") // foo_and_bar
```

<details>
<summary>Examples</summary>

```golang
ToSnakeCase("fooAndBar")   // foo_and_bar
ToSnakeCase("foo.AndBar")  // foo_and_bar
ToSnakeCase("Foo-And-Bar") // foo_and_bar
ToSnakeCase("foo And Bar") // foo_and_bar
ToSnakeCase("你好")         // 你_好
```

</details>

### ToKebabCase
Convert string to ToKebabCase
```golang
ToKebabCase("fooAndBar") // foo-and-bar
```

<details>
<summary>Examples</summary>

```golang
ToKebabCase("fooAndBar")   // foo-and-bar
ToSnakeCase("foo.AndBar")  // foo-and-bar
ToSnakeCase("Foo-And-Bar") // foo-and-bar
ToSnakeCase("foo And Bar") // foo-and-bar
ToSnakeCase("你好")         // 你-好
```

</details>

### Humanize
Converts a computerized string into a human-friendly one
```golang
Humanize("fooAndBar") // foo and bar
```

<details>
<summary>Examples</summary>

```golang
Humanize("fooAndBar")    // foo and bar
Humanize("foo_and_bar")  // foo and bar
Humanize(" foo-and-bar") // foo and bar
```

</details>

### Truncate
Truncates a string up to a specified string length
```golang
Truncate("King Gizzard & The Lizard Wizard", 15) // "King Gizzard..."
```

<details>
<summary>Examples</summary>

```golang
Truncate("En un lugar de la Mancha, de cuyo nombre no quiero acordarme", 27) // "En un lugar de la Mancha..."
```

</details>

### IsEmptyOrBlank
Returns true if string is empty or blank
```golang
IsEmptyOrBlank("") // true
```

<details>
<summary>Examples</summary>

```golang
IsEmptyOrBlank("") // true
IsEmptyOrBlank(" ") // true
IsEmptyOrBlank("my string") // false
```

</details>

### IsBlank
Returns true if string is blank
```golang
IsBlank("") // true
```

<details>
<summary>Examples</summary>

```golang
IsBlank("") // true
IsBlank(" ") // true
IsBlank("my string") // false
```

</details>

### IsEmpty
Returns true if string is empty
```golang
IsEmpty("") // true
```

<details>
<summary>Examples</summary>

```golang
IsEmpty("") // true
IsEmpty(" ") // false
IsEmpty("my string") // false
```

</details>

### DeleteWhitespace
Deletes all whitespaces from a String
```golang
DeleteWhitespace("foo and bar") // fooandbar
```

<details>
<summary>Examples</summary>

```golang
DeleteWhitespace("my string") // mystring
DeleteWhitespace(" my string ") // mystring
```

</details>

### CaseInsensitiveEquals
Return true if both strings are case insensitive equal
```golang
CaseInsensitiveEquals("my string", "My String") // true
```

<details>
<summary>Examples</summary>

```golang
CaseInsensitiveEquals("my string", "My String") // true
CaseInsensitiveEquals("my string", "My-String") // false
```

</details>

### IsBoolean
Determines whether the string is a boolean representation
```golang
IsBoolean("true") // true
```

<details>
<summary>Examples</summary>

```golang
IsBoolean("false") // true
IsBoolean("I am a dog") // false
```

</details>

### IsInteger
Determines whether the string is a integer representation
```golang
IsInteger("1") // true
```

<details>
<summary>Examples</summary>

```golang
IsInteger("344") // true
IsInteger("I am a cat") // false
```

</details>

### IsFloat
Determines whether the string is a float representation
```golang
IsFloat("1.4") // true
```

<details>
<summary>Examples</summary>

```golang
IsFloat("34.4") // true
IsFloat("nothing") // false
```

</details>

### IndexOf
Returns index of the specified character or substring in a particular String
```golang
IndexOf("dog", "o") // 1
```

<details>
<summary>Examples</summary>

```golang
IndexOf("dog", "o") // 1
IndexOf("airport", "po") // 3
```

</details>

## Slice

Import
```go
import "github.com/fegoa89/goutils/slice"
```

### SlicePush

Adds an elements to the end of a slice and returns the new length of the slice.
```golang
mySlice := []interface{}{}
SlicePush(&mySlice, "hello") // 1
```

### SlicePop

Removes the last element from an slice and returns that removed element
```golang
mySlice := []interface{}{"red", "fox"}
SlicePop(&mySlice) // "fox"
```

### SliceShift

Removes the first element from an slice and returns that removed element.
```golang
mySlice := []interface{}{"red", "fox"}
SliceShift(&mySlice) // "red"
```

### RemoveItemByIndex

Removes an item by index position and returns the new length of the slice.
```golang
mySlice := []interface{}{"red", "fox"}
RemoveItemByIndex(&mySlice, 0) // "1"
```

### SliceUnique

Assigns to the slice a duplicate-free version of itself.
```golang
mySlice := []interface{}{"red", "red", "fox"}
SliceUnique(&mySlice) // "1"
```

### SliceDiff

Returns a slice containing all the entries from slice1 that are not present in slice2.
```golang
firstSlice := []interface{}{"red", "dog", "fox"}
secondSlice := []interface{}{"red", "fox"}

SliceDiff(&firstSlice, &secondSlice) // []interface{}{"dog"}
```

### SliceIntersect

Returns a slice containing all the entries from slice1 that are present in slice2.
```golang
firstSlice := []interface{}{"red", "dog", "fox"}
secondSlice := []interface{}{"dog"}

SliceIntersect(&firstSlice, &secondSlice) // []interface{}{"dog"}
```

<br>[⬆ Back to top](#table-of-contents)
