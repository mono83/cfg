package cfg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsErrKeyMissing(t *testing.T) {
	a := assert.New(t)

	a.True(IsErrKeyMissing(ErrKeyMissing{}))
	a.True(IsErrKeyMissing(ErrKeyMissing{Key: "any"}))
	a.False(IsErrKeyMissing(fmt.Errorf("any")))
	a.False(IsErrKeyMissing(nil))
}
