package main

import (
	"fmt"
	"github.com/headwindfly/clevergo"
	"github.com/valyala/fasthttp"
)

// First Middleware
type firstMiddleware struct {
}

func (fm *firstMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		fmt.Fprint(ctx, "I am First Middleware!\n")
		// Invoke the next middleware
		next.Handle(ctx)
	})
}

// Second Middleware
type secondMiddleware struct {
}

func (sm *secondMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		fmt.Fprint(ctx, "I am Second Middleware!\n")
		// Invoke the next middleware
		next.Handle(ctx)
	})
}

// Filter Middleware
type filterMiddleware struct {
}

func (fm *filterMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		// If the name is equal to headwindfly, the request will be blocked.
		if name := ctx.FormValue("name"); len(name) > 0 && string(name) == "headwindfly" {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			fmt.Fprint(ctx, "You are not allow to visit this page!\n")
			return
		}
		// Invoke the next middleware
		next.Handle(ctx)
	})
}

func middleware(ctx *clevergo.Context) {
	ctx.HTML(`<html><head>Middleware Example<title></title></head>
	<body>
		<h2>Hello World.</h2>
		<p><a target="_blank" href="http://localhost:8080?name=headwindfly">http://localhost:8080?name=headwindfly</a> will be blocked by the FilterMiddleware.</p>
		<p><a href="http://localhost:8080?name=othername">http://localhost:8080?name=othername</a> will not be blocked by the FilterMiddleware.</p>
	</body>
	</html>`)
}

func main() {
	app := clevergo.NewApplication()

	// Create a router instance.
	router := app.NewRouter("")

	// Add middleware before registering route's handler.
	router.AddMiddleware(&firstMiddleware{})
	router.AddMiddleware(&secondMiddleware{})

	// If the name is equal to headwindfly, the request will be blocked.
	router.AddMiddleware(&filterMiddleware{})

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(middleware))

	// Start server.
	app.Run()
}
