package panic

import (
	"runtime"
	"strings"
	"fmt"
	"errors"
)

// Taken from github gist: https://gist.github.com/swdunlop/9629168
// Pulled into a package so i could reuse
func IdentifyPanic() string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	switch {
	case name != "":
		return fmt.Sprintf("%v:%v", name, line)
	case file != "":
		return fmt.Sprintf("%v:%v", file, line)
	}

	return fmt.Sprintf("pc:%x", pc)
}

// ReturnPanic
// @param rErr *error  A pointer to an error variable that will contain the panic
//
// ReturnPanic is meant to be deferred in a function that might panic. It will recover
// the panic into an error, which it will then set into the variable specified in a pointer.
// This allows your function to handle the panic and return it like a normal err
//
// Example:
// func Foo(i int) (j int, err error) {
//     defer ReturnPanic(&err)
//     i2, err := funcThatWontPanic(i)
//     j := funcThatMightPanic(i2)
//
//     return
// }
func ReturnPanic(rErr *error) {
	if r := recover(); r != nil {
		// Coerce the panic into an error
		switch x := r.(type) {
		case string:
			*rErr = errors.New(x)
		case error:
			*rErr = x
		default:
			*rErr = errors.New("Unknown panic")
		}
	}
}