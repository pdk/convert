package convert

import (
	"fmt"
	"reflect"
)

// ToInterfaces converts its input into a []interface{}. Returns a new slice,
// true if successful. Returns nil, false if input is not a slice.
func ToInterfaces(in interface{}) ([]interface{}, bool) {

	slice := reflect.ValueOf(in)
	if slice.Kind() != reflect.Slice {
		return nil, false
	}

	l := slice.Len()
	out := make([]interface{}, l)

	for i := 0; i < l; i++ {
		out[i] = slice.Index(i).Interface()
	}

	return out, true
}

// ToStrings takes a slice of any type, and returns a []string, built by calling
// .String() on each item. Returns nil, false if the input is not a slice, or if
// any item in the slice does not have a .String() method. Returns a new slice
// of strings, true if successful.
func ToStrings(in interface{}) ([]string, bool) {

	strings, ok := ToInterfaces(in)
	if !ok {
		return nil, false
	}

	out := make([]string, 0, len(strings))

	for _, i := range strings {
		stringer, ok := i.(fmt.Stringer)
		if !ok {
			return nil, false
		}
		out = append(out, stringer.String())
	}

	return out, true
}
