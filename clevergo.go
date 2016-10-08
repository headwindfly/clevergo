// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package clevergo

import (
	"fmt"
	"time"
)

const (
	// Version of CleverGo.
	version = "2.0.0"

	// Logo of CleverGo.
	logo = `  ____ _     _______     _______ ____   ____  ___
 / ___| |   | ____\ \   / / ____|  _ \ / ___|/ _ \
| |   | |   |  _|  \ \ / /|  _| | |_) | |  _| | | |
| |___| |___| |___  \ V / | |___|  _ <| |_| | |_| |
 \____|_____|_____|  \_/  |_____|_| \_\\____|\___/ `
)

// Version returns current version of CleverGo.
func Version() string {
	return version
}

func info() {
	fmt.Printf("\x1b[36;1m%s %s\x1b[0m\n\n\x1b[32;1mStarted at %s\x1b[0m\n", logo, version, time.Now())
}

// Middleware Interface.
type Middleware interface {
	Handle(next Handler) Handler // handle request.
}

// A Handler responds to an HTTP request.
type Handler interface {
	Handle(*Context)
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.
type HandlerFunc func(*Context)

// Handle calls f(ctx).
func (f HandlerFunc) Handle(ctx *Context) {
	f(ctx)
}
