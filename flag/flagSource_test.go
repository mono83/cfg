package flag

import (
	"flag"
	"github.com/mono83/cfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlagSource(t *testing.T) {
	a := assert.New(t)

	args := []string{"-v", "-q", "--no-ansi", "--domain=local", "-port=33"}
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.Bool("v", false, "")
	fs.Bool("q", false, "")
	fs.Bool("no-ansi", false, "")
	fs.String("domain", "", "")
	fs.Int("port", 2, "")

	f := NewCustomFlagSource(fs, args)
	fv, _ := f.(cfg.Validable)
	a.NoError(fv.Validate())
	a.True(f.Has("v"))
	a.True(f.Has("q"))
	a.True(f.Has("no-ansi"))
	a.True(f.Has("domain"))
	a.True(f.Has("port"))

	b := false
	a.NoError(f.UnmarshallKey("v", &b))
	a.True(b)

	b = false
	a.NoError(f.UnmarshallKey("q", &b))
	a.True(b)

	b = false
	a.NoError(f.UnmarshallKey("no-ansi", &b))
	a.True(b)

	s := ""
	a.NoError(f.UnmarshallKey("domain", &s))
	a.Equal("local", s)

	i := 0
	a.NoError(f.UnmarshallKey("port", &i))
	a.Equal(33, i)
}
