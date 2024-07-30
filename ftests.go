// Package ftests provides a simple and systematic way to test functions:
// we merely check that on given inputs, functions indeed compute the
// expected outputs.
package ftests

// https://tales.mbivert.com/on-a-function-based-test-framework/

import (
	"encoding/json" // pretty-printing
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type Test struct {
	Name     string
	Fun      interface{}
	Args     []interface{}
	Expected []interface{}
}

// Compute the filename where the given function has been defined.
func getFn(f interface{}) string {
	xs := strings.Split((runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()), ".")
	return xs[len(xs)-1]
}

// Test that the function (f) applied to the given input (args)
// returns the expected output (expected). Calls t.Fatalf() on failure.
func Run1(t *testing.T, f interface{}, args []interface{}, expected []interface{}) {
	// []interface{} -> []reflect.Value
	var vargs []reflect.Value
	for _, v := range args {
		vargs = append(vargs, reflect.ValueOf(v))
	}

	got := reflect.ValueOf(f).Call(vargs)

	// []reflect.Value -> []interface{}
	var igot []interface{}
	for _, v := range got {
		igot = append(igot, v.Interface())
	}

	if !reflect.DeepEqual(igot, expected) {
		sgot, err := json.MarshalIndent(igot, "", "\t")
		if err != nil {
			sgot = []byte(fmt.Sprintf("%+v (%s)", igot, err))
		}
		sexp, err := json.MarshalIndent(expected, "", "\t")
		if err != nil {
			sexp = []byte(fmt.Sprintf("%+v (%s)", expected, err))
		}
		// >= 4 and we get nothing; 3 is asm, 2 is testing, 1 is ftests.Run()
		// not sure we can do better
		/*
			_, fn, l, ok := runtime.Caller(3)
			if !ok {
				fn = "???"
				l = 0
			}
			fmt.Printf("%s:%d got: '%s', expected: '%s'", fn, l, igot, expected)
		*/
		// meh, error are printed as {} with JSON.
		fmt.Printf("got: '%s', expected: '%s'", igot, expected)
		t.Fatalf("got: '%s', expected: '%s'", sgot, sexp)
	}
}

// Perform ftests.Run1() on all the given tests. Stop on first
// failure.
func Run(t *testing.T, tests []Test) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s()/%s", getFn(test.Fun), test.Name), func(t *testing.T) {
			Run1(t, test.Fun, test.Args, test.Expected)
		})
	}
}
