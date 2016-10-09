package clevergo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/clevergo/router"
	"github.com/clevergo/sessions"
	"github.com/valyala/fasthttp"
	"html/template"
	"sync"
)

var contextPool = &sync.Pool{
	New: func() interface{} {
		return &Context{}
	},
}

// Context of request.
//
// It contains the router, session and params.
type Context struct {
	router *Router
	*fasthttp.RequestCtx
	RouterParams *router.Params
	Session      *sessions.Session
}

// NewContext returns a Context instance.
//
// Firstly, it will try to get Context instance from contextPool.
// If failed to get Context from contextPool,
// returns a new Context instance.
func NewContext(r *Router, ctx *fasthttp.RequestCtx, rps *router.Params) *Context {
	if context, ok := contextPool.Get().(*Context); ok {
		context.router = r
		context.RequestCtx = ctx
		context.RouterParams = rps
		return context
	}

	return &Context{
		router:       r,
		RequestCtx:   ctx,
		RouterParams: rps,
	}
}

// Close Context.
//
// Context should be closed after finishing request,
// and at this moment, put the context into contextPool.
func (ctx *Context) Close() {
	ctx.Session = nil
	contextPool.Put(ctx)
}

// SessionStore returns the session store of router.
func (ctx *Context) SessionStore() sessions.Store {
	return ctx.router.sessionStore
}

// Logger returns logger.
//
// Returns the router's logger if the logger is non-nil.
// Otherwise, returns the default logger of ctx.
func (ctx *Context) Logger() fasthttp.Logger {
	if ctx.router.logger != nil {
		return ctx.router.logger
	}
	return ctx.RequestCtx.Logger()
}

const (
	// ContentTypeHTML HTML's ContentType
	ContentTypeHTML = "text/html; charset=utf-8"
	// ContentTypeJSON JSON's ContentType
	ContentTypeJSON = "application/json; charset=utf-8"
	// ContentTypeJSONP JSONP's ContentType
	ContentTypeJSONP = "application/javascript; charset=utf-8"
	// ContentTypeXML XML's ContentType
	ContentTypeXML = "application/xml; charset=utf-8"
)

// SetContentTypeToHTML set Content-Type to HTML.
func (ctx *Context) SetContentTypeToHTML() {
	ctx.Response.Header.Set("Content-Type", ContentTypeHTML)
}

// SetContentTypeToJSON set Content-Type to JSON.
func (ctx *Context) SetContentTypeToJSON() {
	ctx.Response.Header.Set("Content-Type", ContentTypeJSON)
}

// SetContentTypeToJSONP set Content-Type to JSONP.
func (ctx *Context) SetContentTypeToJSONP() {
	ctx.Response.Header.Set("Content-Type", ContentTypeJSONP)
}

// SetContentTypeToXML set Content-Type to XML.
func (ctx *Context) SetContentTypeToXML() {
	ctx.Response.Header.Set("Content-Type", ContentTypeXML)
}

// JSON responses JSON data to client.
func (ctx *Context) JSON(v interface{}) {
	json, err := json.Marshal(v)
	if err != nil {
		fmt.Fprint(ctx, err.Error())
		return
	}
	ctx.SetContentTypeToJSON()
	ctx.Response.SetBody(json)
}

// JSONWithCode responses JSON data and custom status code to client.
func (ctx *Context) JSONWithCode(code int, v interface{}) {
	ctx.Response.SetStatusCode(code)
	ctx.JSON(v)
}

// JSONP responses JSONP data to client.
func (ctx *Context) JSONP(v interface{}, callback []byte) {
	json, err := json.Marshal(v)
	if err != nil {
		fmt.Fprint(ctx, err.Error())
		return
	}
	ctx.SetContentTypeToJSONP()
	jsonp := append(callback, "("...)
	jsonp = append(jsonp, json...)
	jsonp = append(jsonp, ")"...)
	ctx.Response.SetBody(jsonp)
}

// JSONPWithCode responses JSONP data and custom status code to client.
func (ctx *Context) JSONPWithCode(code int, v interface{}, callback []byte) {
	ctx.Response.SetStatusCode(code)
	ctx.JSONP(v, callback)
}

// XML responses XML data to client.
func (ctx *Context) XML(v interface{}, headers ...string) {
	xmlBytes, err := xml.MarshalIndent(v, "", `   `)
	if err != nil {
		fmt.Fprint(ctx, err.Error())
		return
	}

	header := xml.Header
	if len(headers) > 0 {
		header = headers[0]
	}

	var bytes []byte
	bytes = append(bytes, header...)
	bytes = append(bytes, xmlBytes...)

	ctx.SetContentTypeToXML()
	ctx.Response.SetBody(bytes)
}

// XMLWithCode responses XML data and custom status code to client.
func (ctx *Context) XMLWithCode(code int, v interface{}, headers ...string) {
	ctx.Response.SetStatusCode(code)
	ctx.XML(v, headers...)
}

// HTML responses HTML data to client.
func (ctx *Context) HTML(body string) {
	ctx.SetContentTypeToHTML()
	ctx.Response.SetBodyString(body)
}

// HTMLWithCode responses HTML data and custom status code to client.
func (ctx *Context) HTMLWithCode(code int, body string) {
	ctx.Response.SetStatusCode(code)
	ctx.HTML(body)
}

// Text responses text data to client using fmt.Fprint().
func (ctx *Context) Text(a ...interface{}) {
	fmt.Fprint(ctx, a...)
}

// Textf responses text data to client using fmt.Fprintf().
func (ctx *Context) Textf(format string, a ...interface{}) {
	fmt.Fprintf(ctx, format, a...)
}

// Render for rendering a template.
func (ctx *Context) Render(tpl *template.Template, data interface{}) {
	ctx.SetContentTypeToHTML()
	tpl.Execute(ctx, data)
}
