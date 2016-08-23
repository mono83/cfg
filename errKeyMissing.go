package cfg

import "fmt"

// ErrKeyMissing is special error, thrown when missing key is requested
type ErrKeyMissing struct {
	Key string
}

func (e ErrKeyMissing) Error() string {
	return fmt.Sprintf("Key with name %s not found in config", e.Key)
}

// IsErrKeyMissing returns true if provided error is instance of ErrKeyMissing
func IsErrKeyMissing(e error) bool {
	_, ok := e.(ErrKeyMissing)
	return ok
}
