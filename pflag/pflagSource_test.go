package pflag

import (
	"github.com/mono83/cfg"
	"github.com/ogier/pflag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlagSource(t *testing.T) {
	a := assert.New(t)

	args := []string{"-v", "arg", "-q", "--no-ansi", "--domain=local", "--port=33", "arg"}
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	fs.BoolP("verbose", "v", false, "")
	fs.BoolP("quiet", "q", false, "")
	fs.Bool("no-ansi", false, "")
	fs.String("domain", "", "")
	fs.Int("port", 2, "")

	f := NewCustomPFlagSource(fs, args)
	fv, _ := f.(cfg.Validable)
	a.NoError(fv.Validate())
	a.True(f.Has("verbose"))
	a.True(f.Has("quiet"))
	a.True(f.Has("no-ansi"))
	a.True(f.Has("domain"))
	a.True(f.Has("port"))

	b := false
	a.NoError(f.UnmarshalKey("verbose", &b))
	a.True(b)

	b = false
	a.NoError(f.UnmarshalKey("quiet", &b))
	a.True(b)

	b = false
	a.NoError(f.UnmarshalKey("no-ansi", &b))
	a.True(b)

	s := ""
	a.NoError(f.UnmarshalKey("domain", &s))
	a.Equal("local", s)

	i := 0
	a.NoError(f.UnmarshalKey("port", &i))
	a.Equal(33, i)
}
