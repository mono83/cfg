package inject

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGraph(t *testing.T) {
	assert := assert.New(t)

	graph := NewGraph(nil)
	assert.NotNil(graph)

	assert.False(graph.HasService("math"))
	buildInvoked := false
	graph.Define("math", func(Container) (interface{}, error) {
		buildInvoked = true
		return func(x, y int) int {
			return x + y
		}, nil
	})
	assert.True(graph.HasService("math"))
	assert.False(buildInvoked)
	graph.Define("printer", func(i Container) (interface{}, error) {
		var f func(int, int) int
		err := graph.GetService("math", &f)
		if err != nil {
			return nil, err
		}

		return func(x, y int) string {
			z := f(x, y)
			return strconv.Itoa(z)
		}, nil
	})
	assert.False(buildInvoked)
	var f func(int, int) string
	err := graph.GetService("printer", &f)
	assert.NoError(err)
	assert.True(buildInvoked)
	assert.Equal("8", f(3, 5))

	// No panic
	var f2 func(int) string
	err = graph.GetService("printer", &f2)
	assert.Error(err)
	assert.Equal("Error copying service printer - reflect.Set: value of type func(int, int) string is not assignable to type func(int) string", err.Error())
}
