# godi
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
	"github.com/neko-neko/godi"
	"fmt"
)

type Dependency struct {
	Num int
	Message string
}

type InjectTarget struct {
	Dep *Dependency `inject:""`
}

func main() {
	var dependency *Dependency = &Dependency{
		Num: 100,
		Message: "Hello depends",
	}
	var target *InjectTarget = &InjectTarget{}

	inj := inject.NewInjector()
	inj.Provide(dependency)
	err := inj.Inject(target)
	if err != nil {
		panic(err)
	}

	fmt.Println(target.Dep.Num)
	fmt.Println(target.Dep.Message)
}
```
