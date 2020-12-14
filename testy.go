package testy

import (
	"reflect"
	"strings"
	"testing"
)

// notErr() - Check to see if the returned err is NOT nil
func notErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

// isErr()  -  Should return an error condition
func isErr(t *testing.T, err error, msg string) {
	if err == nil {
		t.Error(msg)
	}
}

// assert() if the two values are equal, if not error out
func assert(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received \"%v\" (type %v), expected \"%v\" (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

// assertNot() that the two values are NOT equal
func assertNot(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		return
	}
	t.Errorf("Received \"%v\" (type %v), but NOT expecting to", a, reflect.TypeOf(a))
}

// contains() - does the string contain somewhere the value?
func contains(t *testing.T, src string, v string) {
	if strings.Contains(src, v) {
		return
	}
	t.Errorf("String \"%s\" does not Contain \"%s\"", src, v)
}

// arrContains() - does the []string contain the value?
func arrContains(t *testing.T, src []string, v string) {
	for i := range src {
		if src[i] == v {
			return
		}
	}
	t.Errorf("Slice \"%v\" does not Contain \"%s\"", src, v)
}
