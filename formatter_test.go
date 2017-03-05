package main

import (
	"testing"
)

type StubInfo struct{}

func (info StubInfo) Path() string {
	return "path"
}
func (into StubInfo) Size() int64 {
	return 10
}

func TestFormat(t *testing.T) {
	f := DefaultFormatter{}
	assertEqual(t, f.format("%p %s", StubInfo{}), "path 10")
	assertEqual(t, f.format("%p,%s", StubInfo{}), "path,10")
	assertEqual(t, f.format("%s%p", StubInfo{}), "10path")
}
