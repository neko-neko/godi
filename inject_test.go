package inject

import (
	"fmt"
	"reflect"
	"testing"
)

// TestLogger logger mock
type TestLogger struct {
}

func (l *TestLogger) Debugf(format string, v ...interface{}) {
}

// TestDependency1 is depends object
type TestDependency1 struct {
	Num     int
	Message string
}

// TestDependency2 is depends object
type TestDependency2 struct {
	Num     int
	Message string
}

// TestInjectTarget1 is injection target object
type TestInjectTarget1 struct {
	Dep *TestDependency1 `inject:""`
}

// TestInjectTarget2 is injection target object
type TestInjectTarget2 struct {
	Dep TestDependency1 `inject:""`
}

// TestInjectTarget3 is injection target object
type TestInjectTarget3 struct {
	Dep int `inject:""`
}

// TestInjectTarget4 is injection target object
type TestInjectTarget4 struct {
	Dep *int `inject:""`
}

// TestNewInjector verify return not nil?
func TestNewInjector(t *testing.T) {
	inj := NewInjector()
	if inj == nil {
		t.Error("NewContainer return nil")
	}
}

// TestNewInjectorWithLogger verify return not nil and contains Logger?
func TestNewInjectorWithLogger(t *testing.T) {
	inj := NewInjectorWithLogger(&TestLogger{})
	if inj.debugLogger == nil {
		t.Error("NewInjectorWithLogger logger return nil")
	}
}

// TestInjectorProvide verify put container a object?
func TestInjectorProvideObjectIntoContainer(t *testing.T) {
	inj := NewInjector()
	var dependency *TestDependency1 = &TestDependency1{
		Num:     100,
		Message: "Hello inject",
	}
	inj.Provide(dependency)
	containerObj := inj.container.Get(reflect.TypeOf(dependency))

	if reflect.ValueOf(dependency) != containerObj {
		t.Error("Provide put invalid value")
	}
}

// TestInjectorProvide verify put container objects?
func TestInjectorProvideObjectsIntoContainer(t *testing.T) {
	inj := NewInjector()
	var dependency1 *TestDependency1 = &TestDependency1{
		Num:     1,
		Message: "Hello",
	}
	var dependency2 *TestDependency2 = &TestDependency2{
		Num:     2,
		Message: "Inject",
	}
	inj.Provide(dependency1, dependency2)

	containerObj1 := inj.container.Get(reflect.TypeOf(dependency1))
	if reflect.ValueOf(dependency1) != containerObj1 {
		t.Error("Provide put invalid value")
	}

	containerObj2 := inj.container.Get(reflect.TypeOf(dependency2))
	if reflect.ValueOf(dependency2) != containerObj2 {
		t.Error("Provide put invalid value")
	}
}

// TestInjectorProvide verify inject depends
func TestInjectorProvide(t *testing.T) {
	var dependency *TestDependency1 = &TestDependency1{
		Num:     100,
		Message: "Hello inject",
	}
	var target *TestInjectTarget1 = &TestInjectTarget1{}

	inj := NewInjector()
	inj.Provide(dependency)
	err := inj.Inject(target)
	if err != nil {
		t.Error(err)
	}
	if target.Dep.Num != 100 {
		t.Errorf("inject invalid inject value for inj.Dep.Num: %d", target.Dep.Num)
	}
	if target.Dep.Message != "Hello inject" {
		t.Errorf("inject invalid inject value for inj.Dep.Message: %s", target.Dep.Message)
	}
}

// ExampleInjectStructType inject type of struct testing
func ExampleInjectStructType() {
	var dependency TestDependency1 = TestDependency1{
		Num:     100,
		Message: "Hello inject",
	}
	var target *TestInjectTarget2 = &TestInjectTarget2{}

	inj := NewInjector()
	inj.Provide(dependency)
	err := inj.Inject(target)
	if err != nil {
		panic(err)
	}
	fmt.Println(target.Dep.Num)
	fmt.Println(target.Dep.Message)

	// Unordered output:
	// 100
	// Hello inject
}

// ExampleInjectStructPointerType inject type of struct pointer testing
func ExampleInjectStructPointerType() {
	var dependency *TestDependency1 = &TestDependency1{
		Num:     100,
		Message: "Hello inject",
	}
	var target *TestInjectTarget1 = &TestInjectTarget1{}

	inj := NewInjector()
	inj.Provide(dependency)
	err := inj.Inject(target)
	if err != nil {
		panic(err)
	}
	fmt.Println(target.Dep.Num)
	fmt.Println(target.Dep.Message)

	// Unordered output:
	// 100
	// Hello inject
}

// ExampleInjectIntType inject type of int testing
func ExampleInjectIntType() {
	var dependency int = 5
	var target *TestInjectTarget3 = &TestInjectTarget3{}

	inj := NewInjector()
	inj.Provide(dependency)
	err := inj.Inject(target)
	if err != nil {
		panic(err)
	}
	fmt.Println(target.Dep)

	// Unordered output:
	// 5
}

// ExampleInjectIntPointerType inject type of int pointer testing
func ExampleInjectIntPointerType() {
	var num int = 5
	var dependency *int = &num
	var target *TestInjectTarget4 = &TestInjectTarget4{}

	inj := NewInjector()
	inj.Provide(dependency)
	err := inj.Inject(target)
	if err != nil {
		panic(err)
	}
	fmt.Println(*target.Dep)

	// Unordered output:
	// 5
}
