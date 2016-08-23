package cfg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type validableEntry struct {
	id int
}

func (v validableEntry) Validate() error {
	if v.id != 5 {
		return fmt.Errorf("Only 5 is allowed, received %d", v.id)
	}

	return nil
}

func TestValidableConfig(t *testing.T) {
	a := assert.New(t)

	sm := map[string]interface{}{
		"foo": validableEntry{id: 5},
		"bar": validableEntry{id: 1},
	}

	c := NewValidableConfig(Map(sm))

	a.True(c.Has("foo"))
	a.False(c.Has("Foo"))
	a.True(c.Has("bar"))

	v := validableEntry{}
	a.NoError(c.UnmarshallKey("foo", &v))
	a.Equal(5, v.id)
	a.Error(c.UnmarshallKey("bar", &v))
	a.Equal(1, v.id)
}
