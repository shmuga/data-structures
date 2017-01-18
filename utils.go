package data_structures

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func getFatal(t *testing.T, msg interface{}, a interface{}, b interface{}, file string, no int, ok bool) {
	if ok {
		arr := strings.Split(file, "/")
		name := arr[len(arr)-1]
		errmsg := fmt.Sprintf("\n%s:%s %v\nExpected=%v\nActual=%v",
			name,
			strconv.Itoa(no),
			msg,
			a,
			b,
		)
		t.Fatal(errmsg)
	} else {
		t.Fatal(msg)
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, msg interface{}) {
	if a != b {
		_, f, no, ok := runtime.Caller(1)
		getFatal(t, msg, a, b, f, no, ok)
	}
}

func assert(t *testing.T, cond bool, msg interface{}) {
	if !cond {
		_, f, no, ok := runtime.Caller(1)
		getFatal(t, msg, true, cond, f, no, ok)
	}
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}, msg interface{}) {
	if a == b {
		_, f, no, ok := runtime.Caller(1)
		getFatal(t, msg, a, b, f, no, ok)
	}
}
