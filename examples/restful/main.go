// RESTful API Controller Example.

package main

import (
	"github.com/headwindfly/clevergo"
	"html/template"
)

var (
	html = `<html>
	<head></head>
	<body>
		<h3>RESTful API Controller Example.</h3>

		<h4>Requests</h4>
		<ul>
			<li><a target="_blank" href="javascript:get();">GET</a></li>
			<li><a target="_blank" href="javascript:post('POST');">POST</a></li>
			<li><a target="_blank" href="javascript:post('DELETE');">DELETE</a></li>
			<li><a target="_blank" href="javascript:post('PUT');">PUT</a></li>
			<li><a target="_blank" href="javascript:post('HEAD');">HEAD (404 NOT FOUND)</a></li>
			<li><a target="_blank" href="javascript:post('OPTIONS');">OPTIONS (404 NOT FOUND)</a></li>
			<li><a target="_blank" href="javascript:post('PATCH');">PATCH (404 NOT FOUND)</a></li>
		</ul>

		<h4>Result:</h4>
		<textarea rows="5" cols="100" id="result"></textarea>
		<br>

		<script>
			var resultEle = document.getElementById("result");

			var get = function(){
				resultEle.value = 'Pending';
				xmlHttp = new XMLHttpRequest();
    				xmlHttp.open("GET", '/users');
    				xmlHttp.send(null);
    				xmlHttp.onreadystatechange = function () {
        				resultEle.value = "GET: " + xmlHttp.responseText;
    				}
			}

			var post = function(type){
				resultEle.value = 'Pending';
				var url = '/users';
				switch(type){
					case 'POST':
						break;
					case 'DELETE':
						url += '?_method=DELETE';
						break;
					case 'PUT':
						url += '?_method=PUT';
						break;
					case 'HEAD':
						url += '?_method=HEAD';
						break;
					case 'OPTIONS':
						url += '?_method=OPTIONS';
						break;
					case 'PATCH':
						url += '?_method=PATCH';
						break;
				}

				xmlHttp = new XMLHttpRequest();
    				xmlHttp.open("POST", url);
				xmlHttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
				xmlHttp.send();
    				xmlHttp.onreadystatechange = function () {
        				resultEle.value = type + ": " + xmlHttp.responseText;
    				}
			}
		</script>
	</body>
	</html>`
	tpl = template.Must(template.New("").Parse(html))
)

// accessControlMiddleware for setting Access-Control-Allow-* into response's header.
type accessControlMiddleware struct {
	origin  string
	methods string
}

// newAccessControlMiddleware returns a accessControlMiddleware's instance.
func newAccessControlMiddleware(origin, methods string) accessControlMiddleware {
	return accessControlMiddleware{
		origin:  origin,
		methods: methods,
	}
}

// Handle implemented the Middleware interface.
func (m accessControlMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		// Set Access-Control-Allow-Origin and Access-Control-Allow-Methods for ajax request.
		ctx.Response.Header.Set("Access-Control-Allow-Origin", m.origin)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", m.methods)

		next.Handle(ctx)
	})
}

type userController struct {
	clevergo.Controller
}

func newUserController(middlewares []clevergo.Middleware) userController {
	return userController{
		Controller: clevergo.Controller{
			Middlewares: middlewares,
		},
	}
}

func simulate(c clevergo.ControllerInterface, next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		// Using param named '_method' to simulate the other request, such as PUT, DELETE etc.
		if ctx.IsPost() {
			switch string(ctx.FormValue("_method")) {
			case "PUT":
				c.PUT(ctx)
				return
			case "DELETE":
				c.DELETE(ctx)
				return
			case "HEAD":
				c.HEAD(ctx)
				return
			case "OPTIONS":
				c.OPTIONS(ctx)
				return
			case "PATCH":
				c.PATCH(ctx)
				return
			}
		}

		next.Handle(ctx)
	})
}

// Handle implemented the Middleware interface.
//
// Important note： your controller have to implement the Middleware interface.
func (c userController) Handle(next clevergo.Handler) clevergo.Handler {
	return simulate(c, next)
}

func (c userController) GET(ctx *clevergo.Context) {
	ctx.Text("GET handler of userController.")
}

func (c userController) POST(ctx *clevergo.Context) {
	ctx.Text("POST handler of userController.")
}

func (c userController) DELETE(ctx *clevergo.Context) {
	ctx.Text("DELETE handler of userController.")
}

func (c userController) PUT(ctx *clevergo.Context) {
	ctx.Text("PUT handler of userController.")
}

func index(ctx *clevergo.Context) {
	ctx.SetContentTypeToHTML()
	tpl.Execute(ctx, nil)
}

func main() {
	app := clevergo.NewApplication()

	// Create a router instance.
	router := app.NewRouter("")

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(index))
	router.RegisterController("/users", newUserController([]clevergo.Middleware{
		newAccessControlMiddleware("*", "GET, POST, DELETE, PUT"),
	}))

	// Start server.
	app.Run()
}
