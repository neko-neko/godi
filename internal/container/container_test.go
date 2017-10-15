package container

import (
	"testing"
	"reflect"
)

// TestNewContainer return not nil?
func TestNewContainer(t *testing.T) {
	c := NewContainer()
	if c == nil {
		t.Error("NewContainer return nil")
	}
}

// TestAddPutItem put item to container?
func TestAddPutItem(t *testing.T) {
	c := NewContainer()

	object := struct{}{}
	key := reflect.TypeOf(object)
	val := reflect.ValueOf(object)
	c.Add(key, val)

	if c[key] != val {
		t.Error("Add stores an invalid value in the container")
	}
}

// TestGetItem return valid value?
func TestGetItem(t *testing.T) {
	c := NewContainer()

	object := struct{}{}
	key := reflect.TypeOf(object)
	val := reflect.ValueOf(object)
	c.Add(key, val)

	if c.Get(key) != val {
		t.Error("Get return invalid value")
	}
}