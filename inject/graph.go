package inject

import (
	"fmt"
	"github.com/mono83/cfg"
	"reflect"
)

type graph struct {
	cfg.Configurer

	buildFunctions map[string]func(Container) (interface{}, error)
	services       map[string]interface{}

	onAfterInitFuncs []func(name string, target interface{}) error
}

// NewGraph builds new graph DI container
func NewGraph(c cfg.Configurer) Container {
	return &graph{
		Configurer:     c,
		services:       map[string]interface{}{},
		buildFunctions: map[string]func(Container) (interface{}, error){},
	}
}

func (g *graph) OnAfterInit(f func(name string, target interface{}) error) {
	if f != nil {
		g.onAfterInitFuncs = append(g.onAfterInitFuncs, f)
	}
}

func (g *graph) HasService(name string) bool {
	_, ok := g.buildFunctions[name]
	if !ok {
		_, ok = g.services[name]
	}

	return ok
}

func (g *graph) getOrBuild(name string) (interface{}, error) {
	if service, ok := g.services[name]; ok {
		return service, nil
	}

	bf, ok := g.buildFunctions[name]
	if ok {
		service, err := bf(g)
		if service == nil && err == nil {
			err = fmt.Errorf("Builder for %s returns nil without error", name)
		}
		if err == nil {
			if len(g.onAfterInitFuncs) > 0 {
				for _, afterInit := range g.onAfterInitFuncs {
					err = afterInit(name, service)
					if err != nil {
						break
					}
				}
			}
			if err == nil {
				g.services[name] = service
			}
		}

		return service, err
	}

	return nil, fmt.Errorf("Service %s not defined", name)
}

func (g *graph) GetService(name string, target interface{}) error {
	service, err := g.getOrBuild(name)
	if err != nil {
		return err
	}

	fg := &failsafeGetter{name: name, source: service, target: target}
	fg.cp()

	return fg.err
}

func (g *graph) MustGetService(name string, target interface{}) {
	err := g.GetService(name, target)
	if err != nil {
		panic(err)
	}
}

func (g *graph) MustGetServices(defs ...Definition) {
	err := g.GetServices(defs...)
	if err != nil {
		panic(err)
	}
}

func (g *graph) GetServices(defs ...Definition) error {
	for _, def := range defs {
		if err := g.GetService(def.Name(), def.Target()); err != nil {
			return err
		}
	}

	return nil
}

func (g *graph) Define(name string, build func(Container) (interface{}, error)) {
	g.buildFunctions[name] = build
}

// failsafeGetter is special component, used to handle panics
// on reflection
type failsafeGetter struct {
	err            error
	name           string
	source, target interface{}
}

func (f *failsafeGetter) cp() {
	defer func() {
		if r := recover(); r != nil {
			f.err = fmt.Errorf("Error copying service %s - %v", f.name, r)
		}
	}()

	s := reflect.TypeOf(f.source)
	t := reflect.TypeOf(f.target)
	if t.Kind() != reflect.Ptr {
		f.err = fmt.Errorf("Target must be pointer, %T provided", f.target)
	} else if s.Kind() == reflect.Ptr && t.Elem().Kind() != reflect.Ptr {
		f.err = fmt.Errorf("Source is pointer, so target must be pointer to pointer, %T provided", f.target)
	} else {
		reflect.ValueOf(f.target).Elem().Set(reflect.ValueOf(f.source))
	}
}
