package inject

// GetServices handles multiple service retrieval request
func GetServices(from Container, defs ...Definition) error {
	for _, def := range defs {
		if err := from.GetService(def.Name(), def.Target()); err != nil {
			return err
		}
	}

	return nil
}

// Definition describes service request definition
type Definition interface {
	Name() string
	Target() interface{}
}

// Def builds new Definition
func Def(name string, target interface{}) Definition {
	return def{
		name:   name,
		target: target,
	}
}

type def struct {
	name   string
	target interface{}
}

func (d def) Name() string        { return d.name }
func (d def) Target() interface{} { return d.target }
