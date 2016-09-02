package reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyHelperSimple(t *testing.T) {
	a := assert.New(t)

	var s string
	src := "some string"
	a.NoError(CopyHelper("foo", src, &s))
	a.Equal(src, s)
	src = "second"
	a.NoError(CopyHelper("foo", &src, &s))
	a.Equal(src, s)

	var i int
	srcI := 10003
	a.NoError(CopyHelper("bar", srcI, &i))
	a.Equal(srcI, i)
	srcI = -345
	a.NoError(CopyHelper("bar", &srcI, &i))
	a.Equal(srcI, i)

	var f float64
	srcF := 0.32423
	a.NoError(CopyHelper("baz", srcF, &f))
	a.Equal(srcI, i)
	srcF = -345
	a.NoError(CopyHelper("bar", &srcF, &f))
	a.Equal(srcF, f)
}

func TestCopyHelperSmartString(t *testing.T) {
	a := assert.New(t)

	b := false
	a.NoError(CopyHelper("foo", "true", &b))
	a.True(b)
	a.NoError(CopyHelper("foo", "false", &b))
	a.False(b)
	a.NoError(CopyHelper("foo", "TRuE", &b))
	a.True(b)
	a.NoError(CopyHelper("foo", "faLSe", &b))
	a.False(b)
	a.NoError(CopyHelper("foo", "yes", &b))
	a.True(b)
	a.NoError(CopyHelper("foo", "no", &b))
	a.False(b)
	a.NoError(CopyHelper("foo", "1", &b))
	a.True(b)
	a.NoError(CopyHelper("foo", "0", &b))
	a.False(b)
}

func TestCopyHelperInterfaceTargetCopy(t *testing.T) {
	a := assert.New(t)

	var target interface{}

	_, err := interfaceTargetCopy("Hello, world", &target)
	a.NoError(err)
	a.Equal("Hello, world", target)
	a.NoError(CopyHelper("foo", "Yahoo", &target))
	a.Equal("Yahoo", target)

	_, err = interfaceTargetCopy(800, &target)
	a.NoError(err)
	a.Equal(800, target)
	a.NoError(CopyHelper("foo", 4234, &target))
	a.Equal(4234, target)

	_, err = interfaceTargetCopy(0.2, &target)
	a.NoError(err)
	a.Equal(0.2, target)
	a.NoError(CopyHelper("foo", 3.123, &target))
	a.Equal(3.123, target)

	var et error
	res, err := interfaceTargetCopy("foo", &et)
	a.NoError(err)
	a.False(res)
}
