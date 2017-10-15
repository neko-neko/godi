package container_test

import (
	"reflect"
	"testing"

	"github.com/neko-neko/godi/internal/container"
)

// TestNewContainer return not nil?
func TestNewContainer(t *testing.T) {
	c := container.NewContainer()
	if c == nil {
		t.Error(`NewContainer() = nil`)
	}
}

// TestAdd put item to container?
func TestAdd(t *testing.T) {
	c := container.NewContainer()

	object := struct{}{}
	key := reflect.TypeOf(object)
	val := reflect.ValueOf(object)
	c.Add(key, val)

	if c[key] != val {
		t.Error(`c.Add(key, val) != val`)
	}
}

// TestGet return valid value?
func TestGet(t *testing.T) {
	c := container.NewContainer()

	object := struct{}{}
	key := reflect.TypeOf(object)
	val := reflect.ValueOf(object)
	c.Add(key, val)

	if c.Get(key) != val {
		t.Error(`c.Get(key) != val`)
	}
}
