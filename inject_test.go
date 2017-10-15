package inject

import (
	"fmt"
	"reflect"
	"testing"
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

// TestNewInjector verify return not nil?
func TestNewInjector(t *testing.T) {
	inj := NewInjector()
	if inj == nil {
		t.Error(`NewInjector() = nil`)
	}
}

// TestNewInjectorWithLogger verify return not nil and contains Logger?
func TestNewInjectorWithLogger(t *testing.T) {
	inj := NewInjectorWithLogger(&ExampleLogger{})
	if inj.debugLogger == nil {
		t.Error(`NewInjectorWithLogger(&TestLogger{}) = nil`)
	}
}

// TestProvide verify put container objects?
func TestProvide(t *testing.T) {
	type TestDependency1 struct {
		Message string
	}
	type TestDependency2 struct {
		Num int
	}
	var dependencies []interface{}
	dependencies = append(dependencies,
		&TestDependency1{
			Message: "Hello inject",
		},
		&TestDependency2{
			Num: 100,
		},
	)

	inj := NewInjector()
	inj.Provide(dependencies...)
	for _, e := range dependencies {
		if reflect.ValueOf(e) != inj.container.Get(reflect.TypeOf(e)) {
			t.Errorf(`inj.Provide(%v) put invalid value(%v)`, reflect.ValueOf(e), inj.container.Get(reflect.TypeOf(e)))
		}
	}
}

// TestInject can set inject dependency?
func TestInject(t *testing.T) {
	type TestDependency struct {
		Message string
	}
	var dependency *TestDependency = &TestDependency{
		Message: "Hello inject",
	}

	type TestInjectTarget struct {
		Dep *TestDependency `inject:""`
	}
	var target *TestInjectTarget = &TestInjectTarget{}

	inj := NewInjector()
	inj.Provide(dependency)
	inj.Inject(target)

	if target.Dep != dependency {
		t.Errorf(`inj.Inject(%v) could not inject dependency(%v)`, target, dependency)
	}
}

// ExampleNewInjectorWithLogger
func ExampleNewInjectorWithLogger() {
	inj := NewInjectorWithLogger(&ExampleLogger{})
	inj.Inject()

	// Output:
	// [godi-debug] call dependency injection
}

// ExampleInject
func ExampleInjector_Inject() {
	inj := NewInjector()
	inj.Provide(&ExampleInterfaceImpl{})

	var target *ExampleTarget = &ExampleTarget{}
	inj.Inject(target)
	target.Dep.Do()

	// Output:
	// Hello example
}
