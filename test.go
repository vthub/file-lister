package main

import (
	"testing"
	"fmt"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Error(fmt.Sprintf("%v != %v", a, b))
	}
}
