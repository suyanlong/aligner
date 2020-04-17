package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	//load("testdata")
	//load("./")

	t.Log(TmpDir())
}

func TestIsDotFile(t *testing.T) {
	assert.False(t, IsDotFile("testdata/c++.cpp"))
	assert.False(t, IsDotFile("testdata"))
	assert.False(t, IsDotFile("/testdata/"))
	assert.False(t, IsDotFile("c++.cpp"))
	assert.True(t, IsDotFile(".c++.cpp"))
	assert.True(t, IsDotFile(".c++.cpp"))
	assert.True(t, IsDotFile("/.testdata/"))
	assert.True(t, IsDotFile("/.testdata"))
	assert.True(t, IsDotFile("/./"))
	assert.True(t, IsDotFile("./"))
}

func TestIsFormatFile(t *testing.T) {
	comment = "#"
	ext = ".py"
	assert.True(t, IsFormatFile("python.py"))

	assert.False(t, IsFormatFile("rust.rss"))
	assert.False(t, IsFormatFile("rust"))
}
