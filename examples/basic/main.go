package main

import (
	"errors"
	"github.com/headwindfly/clevergo"
	"html/template"
	"os"
	"path"
	"strconv"
)

type project struct {
	Name    string `json:"name" xml:"name"`
	Version string `json:"version" xml:"version"`
}

var (
	helloString   = "Hello CleverGo!"
	resourcesPath = path.Join(os.Getenv("GOPATH"), "src", "github.com", "headwindfly", "clevergo")

	p = project{
		Name:    "CleverGo",
		Version: clevergo.Version,
	}
)

func index(ctx *clevergo.Context) {
	ctx.HTML(`<html>
	<head>
		<title>CleverGo Basic Usage</title>
	</head>
	<body>
		<ul>
			<li><a href="/html" target="_blank">HTML</a></li>
			<li><a href="/html?code=404" target="_blank">HTML With Code: 404</a></li>

			<li><a href="/render" target="_blank">Render HTML</a></li>

			<li><a href="/json" target="_blank">JSON</a></li>
			<li><a href="/json?code=404" target="_blank">JSON With Code: 404</a></li>

			<li><a href="/jsonp?callback=callback" target="_blank">JSONP</a></li>
			<li><a href="/jsonp?callback=callback&code=404" target="_blank">JSONP With Code: 404</a></li>

			<li><a href="/xml" target="_blank">XML</a></li>
			<li><a href="/xml?code=404" target="_blank">XML With Code: 404</a></li>

			<li><a href="/params/yourname" target="_blank">Param</a></li>
			<li><a href="/multi-params/param-one/param-two" target="_blank">Multiple Params</a></li>

			<li><a href="/resources/README.md" target="_blank">Static Resources</a></li>
		</ul>
	</body>
	</html>`)
}

func getCode(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, errors.New("No code.")
	}
	return strconv.Atoi(string(b))
}

func html(ctx *clevergo.Context) {
	code, err := getCode(ctx.FormValue("code"))
	if err == nil {
		ctx.HTMLWithCode(code, helloString)
		return
	}

	ctx.HTML(helloString)
}

func json(ctx *clevergo.Context) {
	code, err := getCode(ctx.FormValue("code"))
	if err == nil {
		ctx.JSONWithCode(code, p)
		return
	}

	ctx.JSON(p)
}

func jsonp(ctx *clevergo.Context) {
	callback := ctx.FormValue("callback")

	code, err := getCode(ctx.FormValue("code"))
	if err == nil {
		ctx.JSONPWithCode(code, p, callback)
		return
	}

	ctx.JSONP(p, callback)
}

func xml(ctx *clevergo.Context) {
	code, err := getCode(ctx.FormValue("code"))
	if err == nil {
		ctx.XMLWithCode(code, p)
		return
	}

	ctx.XML(p)
}

func render(ctx *clevergo.Context) {
	tpl := template.Must(template.New("").Parse(`<h1>Hello {{ .Name }}</h1>`))
	ctx.Render(tpl, map[string]interface{}{
		"Name": "CleverGo",
	})
}

func params(ctx *clevergo.Context) {
	name := ctx.RouterParams.ByName("name")
	ctx.Textf("Hello %s.", name)
}

func multiParams(ctx *clevergo.Context) {
	param1 := ctx.RouterParams.ByName("param1")
	param2 := ctx.RouterParams.ByName("param2")
	ctx.Textf("Your params is %s and %s", param1, param2)
}

func main() {
	// Create a application.
	app := clevergo.NewApplication()

	// Create a router instance.
	router := app.NewRouter("")

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(index))

	// HTML
	router.GET("/html", clevergo.HandlerFunc(html))

	// Render HTML
	router.GET("/render", clevergo.HandlerFunc(render))

	// JSON
	router.GET("/json", clevergo.HandlerFunc(json))

	// JSONP
	router.GET("/jsonp", clevergo.HandlerFunc(jsonp))

	// XML
	router.GET("/xml", clevergo.HandlerFunc(xml))

	// Router Params.
	// Navigate to http://127.0.0.1:8080/params/yourname.
	router.GET("/params/:name", clevergo.HandlerFunc(params))
	// Navigate to http://127.0.0.1:8080/multi-params/param1/param2.
	router.GET("/multi-params/:param1/:param2", clevergo.HandlerFunc(multiParams))

	// Static resource files.
	// Navigate to http://127.0.0.1:8080/resources/README.md
	router.ServeFiles("/resources/*filepath", resourcesPath)

	// Start server.
	app.Run()
}
