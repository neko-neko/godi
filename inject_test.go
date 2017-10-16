package inject_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/neko-neko/godi"
)

// TestLogger logger mock
type TestLogger struct {
}

// Debugf print debug message
func (t *TestLogger) Debugf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

// TestNewInjectorWithLoggerNotNil verify return not nil and contains Logger?
func TestNewInjectorWithLoggerNotNil(t *testing.T) {
	inj := inject.NewInjectorWithLogger(&TestLogger{})
	if inj.DebugLogger() == nil {
		t.Error(`NewInjectorWithLogger(&TestLogger{}) = nil`)
	}
}

// TestProvidePutItemIntoContainer verify put container objects?
func TestProvidePutItemIntoContainer(t *testing.T) {
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

	inj := inject.NewInjector()
	inj.Provide(dependencies...)
	for _, e := range dependencies {
		if reflect.ValueOf(e) != inj.Container().Get(reflect.TypeOf(e)) {
			t.Errorf(`inj.Provide(%v) put invalid value(%v)`, reflect.ValueOf(e), inj.Container().Get(reflect.TypeOf(e)))
		}
	}
}

// TestInjectSetInjectItem can set inject dependency?
func TestInjectSetInjectItem(t *testing.T) {
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

	inj := inject.NewInjector()
	inj.Provide(dependency)
	err := inj.Inject(target)
	if err != nil {
		t.Fatalf(`inj.Inject(%v) return error(%v)`, target, err)
	}

	if target.Dep != dependency {
		t.Errorf(`inj.Inject(%v) could not inject dependency(%v)`, target, dependency)
	}
}

// TestInjectNoSetInjectItem
func TestInjectNoSetInjectItem(t *testing.T) {
	type TestDependency struct {
		Message string
	}
	var dependency *TestDependency = &TestDependency{
		Message: "Hello inject",
	}

	type TestInjectTarget struct {
		Dep *TestDependency
	}
	var target *TestInjectTarget = &TestInjectTarget{}

	inj := inject.NewInjector()
	inj.Provide(dependency)
	err := inj.Inject(target)
	if err != nil {
		t.Fatalf(`inj.Inject(%v) return error(%v)`, target, err)
	}
	if target.Dep != nil {
		t.Errorf(`inj.Inject(%v) ignored "inject" tag (%v)`, target, dependency)
	}
}

// TestInjectCouldNotSetDependency could not set dependency
func TestInjectCouldNotSetDependency(t *testing.T) {
	type TestDependency struct {
		Message string
	}
	var dependency *TestDependency = &TestDependency{
		Message: "Hello inject",
	}

	type TestInjectTarget struct {
		dep *TestDependency `inject:""`
	}
	var target *TestInjectTarget = &TestInjectTarget{}

	inj := inject.NewInjector()
	inj.Provide(dependency)
	err := inj.Inject(target)
	if err == nil {
		t.Fatalf(`inj.Inject(%v) return no error`, target)
	}
	if target.dep != nil {
		t.Errorf(`inj.Inject(%v) if set private memeber then return no error(%v)`, target, err)
	}
}

// TestInjectCouldNotFoundDependency
func TestInjectCouldNotFoundDependency(t *testing.T) {
	type TestDependency struct {
		Message string
	}
	type TestInjectTarget struct {
		Dep *TestDependency `inject:""`
	}
	var target *TestInjectTarget = &TestInjectTarget{}

	inj := inject.NewInjector()
	err := inj.Inject(target)
	if err == nil {
		t.Fatalf(`inj.Inject(%v) return no error`, target)
	}
	if target.Dep != nil {
		t.Errorf(`inj.Inject(%v) if container has not no match memeber then return no error(%v)`, target, err)
	}
}
