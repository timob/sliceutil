package sliceutil

import (
	"reflect"
)

/*
append to slice: create new element, append to slice and return index of element.
If element is pointer, create new pointed to value and set element.
*/
func Append(slicePtr interface{}) int {
	r_slicePtr := reflect.ValueOf(slicePtr)
	r_slice := reflect.Indirect(r_slicePtr)
	r_newElem := reflect.Indirect(reflect.New(r_slice.Type().Elem()))

	if r_slice.Type().Elem().Kind() == reflect.Ptr {
		r_newElem.Set(reflect.New(r_slice.Type().Elem().Elem()))
	}

	r_newSlice := reflect.Append(r_slice, r_newElem)
	r_slice.Set(r_newSlice)
	return r_slice.Len() - 1
}

