package cfg

// Validable interface describes entities, that are able to self-test
// internal data and return error if something wrong
type Validable interface {
	Validate() error
}
