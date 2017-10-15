// Package container defined dependency injection container
package container

import "reflect"

// Container is dependency container
type Container map[reflect.Type]reflect.Value

// NewContainer initialize
func NewContainer() Container {
	return make(Container)
}

// Add value to container
func (c Container) Add(key reflect.Type, value reflect.Value) {
	c[key] = value
}

// Get container value by key
func (c Container) Get(key reflect.Type) reflect.Value {
	return c[key]
}
