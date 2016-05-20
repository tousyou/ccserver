package util

import (
	"testing"
    "runtime"
)

func Assert(t *testing.T, a interface{}, b interface{}){
    if a != b {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d", file, line)
    //    t.Errorf("%q not Equal %q",a,b)
    }
}
