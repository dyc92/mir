// Copyright 2019 Michael Li <dyc92@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir

import "fmt"

var (
	engine Engine
)

// SetTag set custom mir's struct tag name(eg: mir)
func SetTag(name string) {
	if name != "" {
		tagName = name
	}
}

// SetDefault set default engine for register handler.
func SetDefault(e Engine) {
	if engine != nil {
		panic("mir: Setup called twice for engine")
	}
	engine = e
}

// RegisterDefault use entries's info to register handler to default engine.
// You must call SetDefault(...)  setup a default engine first or return error.
func RegisterDefault(entries ...interface{}) error {
	if engine == nil {
		return fmt.Errorf("setup a default engine instance first then call this function")
	}
	return Register(engine, entries...)
}

// Register use entries's info to register handler to give engine.
func Register(e Engine, entries ...interface{}) error {
	if e == nil {
		return fmt.Errorf("register entiries to a nil engine")
	}
	tagMirs, err := TagMirFrom(entries...)
	if err != nil {
		return err
	}
	return e.Register(tagMirs)
}
