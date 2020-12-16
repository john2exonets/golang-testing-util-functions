package testy

import (
	"reflect"
	"strings"
	"testing"
)

// NotErr - Check to see if the returned err is NOT nil
func NotErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

// IsErr  -  Should return an error condition
func IsErr(t *testing.T, err error, msg string) {
	if err == nil {
		t.Error(msg)
	}
}

// Assert if the two values are equal, if not error out
func Assert(t *testing.T, a interface{}, b interface{}) {
	if reflect.TypeOf(a).Kind() == reflect.Slice { // Slices have no equality
		if reflect.DeepEqual(a, b) {
			return
		}
	} else {
		if a == b {
			return
		}
	}

	t.Errorf("Received \"%v\" (type %v), expected \"%v\" (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

// AssertNot - that the two values are NOT equal
func AssertNot(t *testing.T, a interface{}, b interface{}) {
	if reflect.TypeOf(a).Kind() == reflect.Slice {	// Slices have no equality
		if ! reflect.DeepEqual(a, b) {
			return
		}
	} else {
		if a != b {
			return
		}
	}
	t.Errorf("Received \"%v\" (type %v), but NOT expecting to", a, reflect.TypeOf(a))
}

// Contains - does the list contain the target?
// Based on: https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
func Contains(t *testing.T, target interface{}, list interface{}) {
	if reflect.TypeOf(list).Kind() == reflect.Slice || reflect.TypeOf(list).Kind() == reflect.Array {
		listvalue := reflect.ValueOf(list)
		for i := 0; i < listvalue.Len(); i++ {
			if target == listvalue.Index(i).Interface() {
				return
			}
		}
	}
	if reflect.TypeOf(target).Kind() == reflect.String && reflect.TypeOf(list).Kind() == reflect.String {
		if strings.Contains(list.(string), target.(string)) {
			return
		}
	}
	t.Errorf("List \"%v\" (type %v) does NOT contain \"%v\" (type %v)", list, reflect.TypeOf(list), target, reflect.TypeOf(target))
}
