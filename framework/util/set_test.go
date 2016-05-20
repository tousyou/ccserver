package util

import (
	"testing"
)

func TestBasicSetOps(t * testing.T) {
	s := NewSet()
	Assert(t,s.Contains(1), false)
	s.Add(1)
	Assert(t,s.Len(), 1)
}
