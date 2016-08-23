package cfg

import "fmt"

type validableAdapter struct {
	Configurer
}

// NewValidableConfig returns adapter for provided configurer, that
// runs Validate function on all returning configuration values if
// this value has Validate method
func NewValidableConfig(source Configurer) Configurer {
	return validableAdapter{Configurer: source}
}

func (v validableAdapter) UnmarshallKey(key string, target interface{}) error {
	err := v.Configurer.UnmarshallKey(key, target)
	if err == nil && target != nil {
		if vt, ok := target.(Validable); ok {
			err = vt.Validate()
			if err != nil {
				err = fmt.Errorf("Validation error on key \"%s\": %s", key, err.Error())
			}
		}
	}

	return err
}

func (v validableAdapter) KeyFunc(key string) func(interface{}) error {
	return ExtractUnmarshallFunc(v, key)
}
