package inject_test

import (
	"fmt"

	"github.com/neko-neko/godi"
)

// ExampleLogger logger mock
type ExampleLogger struct {
}

func (e *ExampleLogger) Debugf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

// ExampleInterface
type ExampleInterface interface {
	Do()
}

// ExampleInterfaceImpl
type ExampleInterfaceImpl struct{}

// Do is example impl
func (e *ExampleInterfaceImpl) Do() {
	fmt.Println("Hello example")
}

// ExampleTarget is example inject target
type ExampleTarget struct {
	Dep ExampleInterface `inject:""`
}

// ExampleNewInjectorWithLogger
func ExampleNewInjectorWithLogger() {
	inj := inject.NewInjectorWithLogger(&ExampleLogger{})
	inj.Inject()

	// Output:
	// [godi-debug] call dependency injection
}

// ExampleInject
func ExampleInjector_Inject() {
	inj := inject.NewInjector()
	inj.Provide(&ExampleInterfaceImpl{})

	var target *ExampleTarget = &ExampleTarget{}
	inj.Inject(target)
	target.Dep.Do()

	// Output:
	// Hello example
}
