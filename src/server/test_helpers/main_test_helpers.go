package test_helpers

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/trusz/web-scientist/server/mock"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

// ToUpperCase _
func ToUpperCase(text string) string {
	return strings.ToUpper(text)
}

// ToLowerCase _
func ToLowerCase(text string) string {
	return strings.ToLower(text)
}

// CreateEqualMocks _
func CreateEqualMocks() (reference *mock.Mock, experiment *mock.Mock) {
	reference = mock.New("localhost", "9999", 0)
	experiment = mock.New("localhost", "9998", 10)

	go reference.Start(nil)
	go experiment.Start(nil)

	return
}

// CreateNonEqualMocks _
func CreateNonEqualMocks() (reference *mock.Mock, experiment *mock.Mock) {
	reference = mock.New("localhost", "9999", 0)
	experiment = mock.New("localhost", "9998", 10)

	go reference.Start(nil)
	go experiment.Start(ToUpperCase)

	return
}
