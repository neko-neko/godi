// Package inject is dependency injection operation package
package inject

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/neko-neko/godi/internal/container"
)

// Logger debbug printer
type Logger interface {
	// Debugf print debugging log
	Debugf(format string, v ...interface{})
}

// debugPrefix is debugging message prefix
const debugPrefix = "[godi-debug]"

// Injector is dependency injection operator
type Injector struct {
	// debug logger
	// If you need debugging dependency injection log, Set logger instance
	debugLogger Logger

	// Dependency container
	container container.Container
}

// NewInjector initialize Injector
func NewInjector() *Injector {
	return &Injector{
		container: container.NewContainer(),
	}
}

// NewInjectorWithLogger initialize injector instance with debugging logger
func NewInjectorWithLogger(logger Logger) *Injector {
	return &Injector{
		debugLogger: logger,
		container:   container.NewContainer(),
	}
}

// Provide object to container
func (inj *Injector) Provide(objects ...interface{}) {
	for _, object := range objects {
		inj.container.Add(reflect.TypeOf(object), reflect.ValueOf(object))
	}
}

// Inject dependency
func (inj *Injector) Inject(targets ...interface{}) error {
	inj.debugf("call dependency injection")

	for _, target := range targets {
		t := reflect.TypeOf(target)
		if k := t.Kind(); k != reflect.Ptr {
			return fmt.Errorf("inject can not support type other than reflect.Ptr: %s", k)
		}
		t = t.Elem()
		v := reflect.ValueOf(target).Elem()

		inj.debugf(fmt.Sprintf("target [type: %s] [value: %s]", t.String(), v.String()))

		for i := 0; i < t.NumField(); i++ {
			refType := t.Field(i)
			_, enabled := refType.Tag.Lookup("inject")
			if !enabled {
				continue
			}

			val := v.Field(i)
			if !val.CanSet() {
				return fmt.Errorf("could not set type for: %s", t.Kind())
			}
			dep, err := inj.loadDep(val.Type())
			if err != nil {
				return err
			}

			val.Set(dep)
		}
	}

	return nil
}

// loadDep returns dependency from container
func (inj *Injector) loadDep(refType reflect.Type) (reflect.Value, error) {
	val := inj.container.Get(refType)
	if val.IsValid() {
		return val, nil
	}

	for k, v := range inj.container {
		if k.Implements(refType) {
			return v, nil
		}
	}

	return val, fmt.Errorf("could not found inject for type: %s", refType.Kind())
}

// debugf output debugging log
func (inj *Injector) debugf(format string, i ...interface{}) {
	if inj.debugLogger == nil {
		return
	}

	inj.debugLogger.Debugf(strings.Join([]string{debugPrefix, format}, " "), i...)
}
