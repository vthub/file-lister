package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, f.Format("%p %s", StubInfo{}), "path 10")
	assert.Equal(t, f.Format("%p,%s", StubInfo{}), "path,10")
	assert.Equal(t, f.Format("%s%p", StubInfo{}), "10path")
}
