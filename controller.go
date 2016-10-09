package clevergo

// ControllerInterface contains the methods which the Controller should implements.
//
// In fact, the controller is a middleware.
type ControllerInterface interface {
	Handle(next Handler) Handler // Implemented Middleware Interface.

	initMiddlewares(next Handler) Handler // Init the middlewares.

	DELETE(ctx *Context)  // Request handler for DELETE request.
	GET(ctx *Context)     // Request handler for GET request.
	HEAD(ctx *Context)    // Request handler for HEAD request.
	OPTIONS(ctx *Context) // Request handler for OPTIONS request.
	PATCH(ctx *Context)   // Request handler for PATCH request.
	POST(ctx *Context)    // Request handler for POST request.
	PUT(ctx *Context)     // Request handler for PUT request.
}

// Controller the conventional RESTful API Controller.
//
// Middlewares for the current controller.
// Important note: the Controller just a sample,
// it shows that how to create a highly scalable RESTful controller.
// Example: https://github.com/headwindfly/clevergo/blob/master/examples/restful.
type Controller struct {
	Middlewares []Middleware
}

func (c *Controller) AddMiddleware(m Middleware) {
	c.Middlewares = append(c.Middlewares, m)
}

// initMiddlewares initialize the controller's middleware.
// Make the controller's handler wraped by additional middlewares.
func (c Controller) initMiddlewares(next Handler) Handler {
	var h Handler
	h = c.Handle(next)

	for i := len(c.Middlewares) - 1; i >= 0; i-- {
		h = c.Middlewares[i].Handle(h)
	}

	return h
}

// Handle implemented Middleware Interface.
func (c Controller) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Invoke the request handler.
		next.Handle(ctx)
	})
}

// DELETE for handling the DELETE request.
func (c Controller) DELETE(ctx *Context) {
	ctx.NotFound()
}

// GET for handling the GET request.
func (c Controller) GET(ctx *Context) {
	ctx.NotFound()
}

// HEAD for handling the HEAD request.
func (c Controller) HEAD(ctx *Context) {
	ctx.NotFound()
}

// OPTIONS for handling the OPTIONS request.
func (c Controller) OPTIONS(ctx *Context) {
	ctx.NotFound()
}

// PATCH for handling the PATCH request.
func (c Controller) PATCH(ctx *Context) {
	ctx.NotFound()
}

// POST for handling the POST request.
func (c Controller) POST(ctx *Context) {
	ctx.NotFound()
}

// PUT for handling the PUT request.
func (c Controller) PUT(ctx *Context) {
	ctx.NotFound()
}
