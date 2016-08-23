package file

import (
	"github.com/mono83/cfg/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	c := New([]string{"testFixture.json"}, ioutil.ReadFile, json.NewBytesSource)
	a.NoError(c.Validate())
	a.True(c.Has("id"))
	a.True(c.Has("name"))

	i := 0
	s := ""
	a.NoError(c.UnmarshalKey("id", &i))
	a.NoError(c.UnmarshalKey("name", &s))
	a.Equal(3874, i)
	a.Equal("Yahoo", s)
	a.NoError(c.Reload())
}
