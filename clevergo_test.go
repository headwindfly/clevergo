package clevergo

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/valyala/fasthttp"
	"net"
	"strconv"
	"testing"
	"time"
)

type expect struct {
	rw       *readWriter
	validate func(fasthttp.Response) error
}

func TestInfo(t *testing.T) {
	info()
}

type statusMiddleware struct {
}

func (m statusMiddleware) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		if status := string(ctx.QueryArgs().Peek("status")); status == "0" {
			ctx.SetBodyString("Invalid status")
			return
		}

		next.Handle(ctx)
	})
}

func TestRouter(t *testing.T) {
	router := NewRouter()
	router.SetMiddlewares([]Middleware{simpleMiddleware{}})
	router.AddMiddleware(statusMiddleware{})

	router.GET("/", HandlerFunc(
		func(ctx *Context) {
			ctx.Text("GET")
		}))
	router.POST("/", HandlerFunc(
		func(ctx *Context) {
			ctx.Text("POST")
		}))
	router.PUT("/", HandlerFunc(
		func(ctx *Context) {
			ctx.Text("PUT")
		}))
	router.DELETE("/", HandlerFunc(
		func(ctx *Context) {
			ctx.Text("DELETE")
		}))
	router.HEAD("/", HandlerFunc(
		func(ctx *Context) {
			ctx.Response.Header.Add("Powered By", "CleverGo")
		}))
	router.OPTIONS("/", HandlerFunc(
		func(ctx *Context) {
			ctx.Response.Header.Add("Access-Control-Request-Method", "POST")
		}))
	router.PATCH("/", HandlerFunc(
		func(ctx *Context) {
			ctx.Text("PATCH")
		}))

	s := &fasthttp.Server{
		Handler: router.Handler,
	}

	expects := make([]expect, 0)

	// GET
	rwGet := &readWriter{}
	rwGet.r.WriteString("GET / HTTP/1.1\r\n\r\n")
	expectGet := expect{
		rw: rwGet,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Body(), []byte("GET")) {
				return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), "GET")
			}
			return nil
		},
	}

	// POST
	rwPost := &readWriter{}
	rwPost.r.WriteString("POST / HTTP/1.1\r\n\r\n")
	expectPost := expect{
		rw: rwPost,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Body(), []byte("POST")) {
				return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), "POST")
			}
			return nil
		},
	}

	// PUT
	rwPut := &readWriter{}
	rwPut.r.WriteString("PUT / HTTP/1.1\r\n\r\n")
	expectPut := expect{
		rw: rwPut,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Body(), []byte("PUT")) {
				return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), "PUT")
			}
			return nil
		},
	}

	// GET
	rwDelete := &readWriter{}
	rwDelete.r.WriteString("DELETE / HTTP/1.1\r\n\r\n")
	expectDelete := expect{
		rw: rwDelete,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Body(), []byte("DELETE")) {
				return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), "DELETE")
			}
			return nil
		},
	}

	// PATCH
	rwPatch := &readWriter{}
	rwPatch.r.WriteString("PATCH / HTTP/1.1\r\n\r\n")
	expectPatch := expect{
		rw: rwPatch,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Body(), []byte("PATCH")) {
				return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), "PATCH")
			}
			return nil
		},
	}

	// HEAD
	rwHead := &readWriter{}
	rwHead.r.WriteString("HEAD / HTTP/1.1\r\n\r\n")
	expectHead := expect{
		rw: rwHead,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Header.Peek("Powered By"), []byte("CleverGo")) {
				return fmt.Errorf("Unexpected Powered By %s. Expected %q", resp.Header.Peek("Powered By"), "CleverGo")
			}
			return nil
		},
	}

	// OPTIONS
	rwOption := &readWriter{}
	rwOption.r.WriteString("OPTIONS / HTTP/1.1\r\n\r\n")
	expectOption := expect{
		rw: rwOption,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Header.Peek("Access-Control-Request-Method"), []byte("POST")) {
				return fmt.Errorf("Unexpected Access-Control-Request-Method %s. Expected %q", resp.Header.Peek("Access-Control-Request-Method"), "POST")
			}
			return nil
		},
	}

	expects = append(
		expects,
		expectGet,
		expectDelete,
		expectOption,
		expectHead,
		expectPut,
		expectPost,
		expectPatch,
	)

	// Status middleware.
	rw := &readWriter{}
	rw.r.WriteString("GET /?status=0 HTTP/1.1\r\n\r\n")

	expects = append(expects, expect{
		rw: rw,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Body(), []byte("Invalid status")) {
				return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), "Invalid status")
			}
			return nil
		},
	})

	for _, v := range expects {
		ch := make(chan error)
		go func() {
			ch <- s.ServeConn(v.rw)
		}()

		select {
		case err := <-ch:
			if err != nil {
				t.Fatalf("return error %s", err)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("timeout")
		}

		br := bufio.NewReader(&v.rw.w)
		var resp fasthttp.Response
		if err := resp.Read(br); err != nil {
			t.Fatalf("Unexpected error when reading response: %s", err)
		}
		if resp.Header.StatusCode() != 200 {
			t.Fatalf("Unexpected status code %d. Expected %d", resp.Header.StatusCode(), 200)
		}

		if !bytes.Equal(resp.Header.Peek("Middleware"), []byte("Simple")) {
			t.Fatalf("Unexpected middleware'name %s. Expected %s", resp.Header.Peek("Middleware"), "Simple")
		}

		if err := v.validate(resp); err != nil {
			t.Fatalf(err.Error())
		}
	}
}

func TestApplication(t *testing.T) {
	app := NewApplication()
	r1 := app.NewRouter("")

	r1.GET("/", HandlerFunc(func(ctx *Context) {
		ctx.Textf("Hello world")
	}))

	r2 := NewRouter()
	app.AddRouter("127.1.1.1", r2)
	r2.GET("/", HandlerFunc(func(ctx *Context) {
		ctx.Textf("127.1.1.1")
	}))

	s := &fasthttp.Server{
		Handler: app.getHandler(),
	}

	expects := make([]expect, 0)

	rw1 := &readWriter{}
	rw1.r.WriteString("GET / HTTP/1.1\r\n\r\n")
	expect1 := expect{
		rw: rw1,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Body(), []byte("Hello world")) {
				return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), "Hello world")
			}
			return nil
		},
	}

	rw2 := &readWriter{}
	rw2.r.WriteString("GET http://127.1.1.1/ HTTP/1.1\r\n\r\n")
	expect2 := expect{
		rw: rw2,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Body(), []byte("127.1.1.1")) {
				return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), "127.1.1.1")
			}
			return nil
		},
	}

	expects = append(
		expects,
		expect1,
		expect2,
	)

	for _, v := range expects {
		ch := make(chan error)
		go func() {
			ch <- s.ServeConn(v.rw)
		}()

		select {
		case err := <-ch:
			if err != nil {
				t.Fatalf("return error %s", err)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("timeout")
		}

		br := bufio.NewReader(&v.rw.w)
		var resp fasthttp.Response
		if err := resp.Read(br); err != nil {
			t.Fatalf("Unexpected error when reading response: %s", err)
		}
		if resp.Header.StatusCode() != 200 {
			t.Fatalf("Unexpected status code %d. Expected %d", resp.Header.StatusCode(), 200)
		}

		if err := v.validate(resp); err != nil {
			t.Fatalf(err.Error())
		}
	}
}

type infoForTest struct {
	XMLName xml.Name `xml:"info"`
	Name    string   `xml:"name";json:"name"`
	Version string   `xml:"version";json:"version"`
}

func TestContext(t *testing.T) {
	info := infoForTest{
		Name:    "CleverGo",
		Version: Version(),
	}

	jsonResult, jsonpResult, xmlResult := []byte{}, []byte{}, []byte{}

	if v, err := json.Marshal(info); err != nil {
		jsonResult = []byte(err.Error())
	} else {
		jsonResult = v
	}

	callback := []byte("callback")
	jsonpResult = append(callback, '(')
	jsonpResult = append(jsonpResult, jsonResult...)
	jsonpResult = append(jsonpResult, ')')

	if v, err := xml.Marshal(info); err != nil {
		xmlResult = []byte(err.Error())
	} else {
		xmlResult = v
	}

	xmlResult = append([]byte(xml.Header), xmlResult...)
	xmlResult = bytes.Replace(xmlResult, []byte("\n"), []byte(""), -1)
	xmlResult = bytes.Replace(xmlResult, []byte(" "), []byte(""), -1)

	router := NewRouter()

	router.GET("/:mode", HandlerFunc(func(ctx *Context) {
		code := -1
		if codeNum, err := strconv.Atoi(string(ctx.QueryArgs().Peek("code"))); err == nil {
			code = codeNum
		}

		switch ctx.Params.String("mode") {
		case "html":
			if code > 0 {
				ctx.HTMLWithCode(code, "CleverGo")
				return
			}
			ctx.HTML("CleverGo")
		case "json":
			if code > 0 {
				ctx.JSONWithCode(code, info)
				return
			}
			ctx.JSON(info)
		case "jsonp":
			if code > 0 {
				ctx.JSONPWithCode(code, info, callback)
				return
			}
			ctx.JSONP(info, callback)
		case "xml":
			if code > 0 {
				ctx.XMLWithCode(code, info)
				return
			}
			ctx.XML(info)
		default:
			ctx.Textf("Unknown mode: %s", ctx.Params.String("mode"))
		}
	}))

	s := &fasthttp.Server{
		Handler: router.Handler,
	}

	expects := make([]expect, 0)

	// HTML
	rw1 := &readWriter{}
	rw1.r.WriteString("GET /html HTTP/1.1\r\n\r\n")
	expect1 := expect{
		rw: rw1,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Header.Peek("Content-Type"), []byte(contentTypeHTML)) {
				return fmt.Errorf("Unexpected Content-Type %q. Expected %q", resp.Header.Peek("Content-Type"), contentTypeHTML)
			}

			if !bytes.Equal(resp.Body(), []byte("CleverGo")) {
				return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), "CleverGo")
			}
			return nil
		},
	}

	// HTML with code.
	rw2 := &readWriter{}
	rw2.r.WriteString("GET /html?code=500 HTTP/1.1\r\n\r\n")
	expect2 := expect{
		rw: rw2,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Header.Peek("Content-Type"), []byte(contentTypeHTML)) {
				return fmt.Errorf("Unexpected Content-Type %q. Expected %q", resp.Header.Peek("Content-Type"), contentTypeHTML)
			}

			if resp.StatusCode() != 500 {
				return fmt.Errorf("Unexpected status code %q. Expected %q", resp.StatusCode(), 500)
			}

			if !bytes.Equal(resp.Body(), []byte("CleverGo")) {
				return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), "CleverGo")
			}
			return nil
		},
	}

	// JSON
	rw3 := &readWriter{}
	rw3.r.WriteString("GET /json HTTP/1.1\r\n\r\n")
	expect3 := expect{
		rw: rw3,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Header.Peek("Content-Type"), []byte(contentTypeJSON)) {
				return fmt.Errorf("Unexpected JSON Content-Type %q. Expected %q", resp.Header.Peek("Content-Type"), contentTypeJSON)
			}

			if !bytes.Equal(resp.Body(), jsonResult) {
				return fmt.Errorf("Unexpected JSON body %q. Expected %q", resp.Body(), string(jsonResult))
			}
			return nil
		},
	}

	// JSON with code.
	rw4 := &readWriter{}
	rw4.r.WriteString("GET /json?code=500 HTTP/1.1\r\n\r\n")
	expect4 := expect{
		rw: rw4,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Header.Peek("Content-Type"), []byte(contentTypeJSON)) {
				return fmt.Errorf("Unexpected JSON Content-Type %q. Expected %q", resp.Header.Peek("Content-Type"), contentTypeJSON)
			}

			if resp.StatusCode() != 500 {
				return fmt.Errorf("Unexpected status code %q. Expected %q", resp.StatusCode(), 500)
			}

			if !bytes.Equal(resp.Body(), jsonResult) {
				return fmt.Errorf("Unexpected JSON body %q. Expected %q", resp.Body(), string(jsonResult))
			}
			return nil
		},
	}

	// JSONP
	rw5 := &readWriter{}
	rw5.r.WriteString("GET /jsonp HTTP/1.1\r\n\r\n")
	expect5 := expect{
		rw: rw5,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Header.Peek("Content-Type"), []byte(contentTypeJSONP)) {
				return fmt.Errorf("Unexpected JSON Content-Type %q. Expected %q", resp.Header.Peek("Content-Type"), contentTypeJSONP)
			}

			if !bytes.Equal(resp.Body(), jsonpResult) {
				return fmt.Errorf("Unexpected JSONP body %q. Expected %q", resp.Body(), string(jsonpResult))
			}
			return nil
		},
	}

	// JSONP with code.
	rw6 := &readWriter{}
	rw6.r.WriteString("GET /jsonp?code=500 HTTP/1.1\r\n\r\n")
	expect6 := expect{
		rw: rw6,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Header.Peek("Content-Type"), []byte(contentTypeJSONP)) {
				return fmt.Errorf("Unexpected JSONP Content-Type %q. Expected %q", resp.Header.Peek("Content-Type"), contentTypeJSONP)
			}

			if resp.StatusCode() != 500 {
				return fmt.Errorf("Unexpected status code %q. Expected %q", resp.StatusCode(), 500)
			}

			if !bytes.Equal(resp.Body(), jsonpResult) {
				return fmt.Errorf("Unexpected JSONP body %q. Expected %q", resp.Body(), string(jsonpResult))
			}
			return nil
		},
	}

	// XML
	rw7 := &readWriter{}
	rw7.r.WriteString("GET /xml HTTP/1.1\r\n\r\n")
	expect7 := expect{
		rw: rw7,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Header.Peek("Content-Type"), []byte(contentTypeXML)) {
				return fmt.Errorf("Unexpected XML Content-Type %q. Expected %q", resp.Header.Peek("Content-Type"), contentTypeXML)
			}

			body := bytes.Replace(resp.Body(), []byte("\n"), []byte(""), -1)
			body = bytes.Replace(body, []byte(" "), []byte(""), -1)
			if !bytes.Equal(body, xmlResult) {
				return fmt.Errorf("Unexpected XML body %q. Expected %q", body, string(xmlResult))
			}
			return nil
		},
	}

	// XML with code.
	rw8 := &readWriter{}
	rw8.r.WriteString("GET /xml?code=500 HTTP/1.1\r\n\r\n")
	expect8 := expect{
		rw: rw8,
		validate: func(resp fasthttp.Response) error {
			if !bytes.Equal(resp.Header.Peek("Content-Type"), []byte(contentTypeXML)) {
				return fmt.Errorf("Unexpected XML Content-Type %q. Expected %q", resp.Header.Peek("Content-Type"), contentTypeXML)
			}

			if resp.StatusCode() != 500 {
				return fmt.Errorf("Unexpected XML status code %q. Expected %q", resp.StatusCode(), 500)
			}

			body := bytes.Replace(resp.Body(), []byte("\n"), []byte(""), -1)
			body = bytes.Replace(body, []byte(" "), []byte(""), -1)
			if !bytes.Equal(body, xmlResult) {
				return fmt.Errorf("Unexpected XML body %q. Expected %q", body, string(xmlResult))
			}
			return nil
		},
	}

	expects = append(
		expects,
		expect1,
		expect2,
		expect3,
		expect4,
		expect5,
		expect6,
		expect7,
		expect8,
	)

	for _, v := range expects {
		ch := make(chan error)
		go func() {
			ch <- s.ServeConn(v.rw)
		}()

		select {
		case err := <-ch:
			if err != nil {
				t.Fatalf("return error %s", err)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("timeout")
		}

		br := bufio.NewReader(&v.rw.w)
		var resp fasthttp.Response
		if err := resp.Read(br); err != nil {
			t.Fatalf("Unexpected error when reading response: %s", err)
		}

		if err := v.validate(resp); err != nil {
			t.Fatalf(err.Error())
		}
	}
}

type userController struct {
	Controller
}

func (c *userController) GET(ctx *Context) {
	ctx.Text("GET")
}

func (c *userController) POST(ctx *Context) {
	ctx.Text("POST")
}

func (c *userController) DELETE(ctx *Context) {
	ctx.Text("DELETE")
}

func (c *userController) PUT(ctx *Context) {
	ctx.Text("PUT")
}

func (c *userController) PATCH(ctx *Context) {
	ctx.Text("PATCH")
}

func (c *userController) OPTIONS(ctx *Context) {
	ctx.Text("OPTIONS")
}

func (c *userController) HEAD(ctx *Context) {
	ctx.Text("HEAD")
}

type simpleMiddleware struct {
}

func (m simpleMiddleware) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		ctx.Response.Header.Add("Middleware", "Simple")
		next.Handle(ctx)
	})
}

func TestController(t *testing.T) {
	app := NewApplication()
	r := NewRouter()

	user := &userController{}
	user.AddMiddleware(simpleMiddleware{})

	r.RegisterController("/", user)
	app.AddRouter("", r)

	s := &fasthttp.Server{
		Handler: app.getHandler(),
	}

	expects := make([]expect, 0)
	methods := []string{"GET", "POST", "DELETE", "PUT", "DELETE", "OPTIONS", "PATCH"}
	for i := 0; i < len(methods); i++ {
		method := methods[i]
		rw := &readWriter{}
		rw.r.WriteString(method + " / HTTP/1.1\r\n\r\n")

		expects = append(expects, expect{
			rw: rw,
			validate: func(resp fasthttp.Response) error {
				if !bytes.Equal(resp.Body(), []byte(method)) {
					return fmt.Errorf("Unexpected body %q. Expected %q", resp.Body(), method)
				}
				return nil
			},
		})
	}

	for _, v := range expects {
		ch := make(chan error)
		go func() {
			ch <- s.ServeConn(v.rw)
		}()

		select {
		case err := <-ch:
			if err != nil {
				t.Fatalf("return error %s", err)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("timeout")
		}

		br := bufio.NewReader(&v.rw.w)
		var resp fasthttp.Response
		if err := resp.Read(br); err != nil {
			t.Fatalf("Unexpected error when reading response: %s", err)
		}
		if resp.Header.StatusCode() != 200 {
			t.Fatalf("Unexpected status code %d. Expected %d", resp.Header.StatusCode(), 200)
		}

		if !bytes.Equal(resp.Header.Peek("Middleware"), []byte("Simple")) {
			t.Fatalf("Unexpected middleware'name %s. Expected %s", resp.Header.Peek("Middleware"), "Simple")
		}

		if err := v.validate(resp); err != nil {
			t.Fatalf(err.Error())
		}
	}
}

type readWriter struct {
	net.Conn
	r bytes.Buffer
	w bytes.Buffer
}

var zeroTCPAddr = &net.TCPAddr{
	IP: net.IPv4zero,
}

func (rw *readWriter) Close() error {
	return nil
}

func (rw *readWriter) Read(b []byte) (int, error) {
	return rw.r.Read(b)
}

func (rw *readWriter) Write(b []byte) (int, error) {
	return rw.w.Write(b)
}

func (rw *readWriter) RemoteAddr() net.Addr {
	return zeroTCPAddr
}

func (rw *readWriter) LocalAddr() net.Addr {
	return zeroTCPAddr
}

func (rw *readWriter) SetReadDeadline(t time.Time) error {
	return nil
}

func (rw *readWriter) SetWriteDeadline(t time.Time) error {
	return nil
}
