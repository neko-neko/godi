package inject_test

import "github.com/neko-neko/godi"

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
