package ini

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var iniSrc = `
[Block]
#Comment
key=1
  key2 = 10
	key3= foo =bar
key5=3.22
// Another comment
`

func TestNewBytesSource(t *testing.T) {
	assert := assert.New(t)

	c, e := NewBytesSource([]byte(iniSrc))
	if assert.NoError(e) && assert.NotNil(c) {
		assert.True(c.Has("key"))
		assert.True(c.Has("key2"))
		assert.True(c.Has("key3"))
		assert.False(c.Has("key4"))
		assert.True(c.Has("key5"))

		var i int
		if assert.NoError(c.UnmarshalKey("key", &i)) {
			assert.Equal(1, i)
		}
		if assert.NoError(c.UnmarshalKey("key2", &i)) {
			assert.Equal(10, i)
		}

		var s string
		if assert.NoError(c.UnmarshalKey("key3", &s)) {
			assert.Equal("foo =bar", s)
		}

		var f float32
		if assert.NoError(c.UnmarshalKey("key5", &f)) {
			assert.Equal(float32(3.22), f)
		}
	}
}
