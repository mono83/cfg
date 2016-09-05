package cfg

// errConfigurer is special configurer that will return
// embedded error on any action
type errConfigurer struct {
	error
}

// ErrConfigurer builds configuration source, that will
// return an error on any action and false on Has invocation
func ErrConfigurer(err error) Configurer {
	return errConfigurer{error: err}
}

func (e errConfigurer) Has(string) bool                        { return false }
func (e errConfigurer) UnmarshalKey(string, interface{}) error { return e }
func (e errConfigurer) KeyFunc(string) func(interface{}) error {
	return func(interface{}) error { return e }
}

func (e errConfigurer) Validate() error { return e }
