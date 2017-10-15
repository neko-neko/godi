# godi
[![GoDoc](https://godoc.org/github.com/neko-neko/godi?status.svg)](https://godoc.org/github.com/neko-neko/godi)
[![Build Status](https://travis-ci.org/neko-neko/godi.svg?branch=master)](https://travis-ci.org/neko-neko/godi)
[![codecov](https://codecov.io/gh/neko-neko/godi/branch/master/graph/badge.svg)](https://codecov.io/gh/neko-neko/godi)
[![Go Report Card](https://goreportcard.com/badge/github.com/neko-neko/godi)](https://goreportcard.com/report/github.com/neko-neko/godi)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/neko-neko/godi/master/LICENSE)

## Overview
godi is light weight simple DI container library for golang.  

## Getting Started
### For [dep](https://github.com/golang/dep) users
When your project top dir run this:  
```bash
$ dep ensure -add github.com/neko-neko/godi
```

### Other users
Run this:  
```bash
$ go get github.com/neko-neko/godi
```

## Example
```go
package main

import (
	"fmt"

	"github.com/neko-neko/godi"
)

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

func main() {
	inj := NewInjector()
	inj.Provide(&ExampleInterfaceImpl{})

	var target *ExampleTarget = &ExampleTarget{}
	inj.Inject(target)
	target.Dep.Do()

	// Output:
	// Hello example
}
```
