// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/clevergo/captcha"
	"github.com/headwindfly/clevergo"
	"github.com/valyala/fasthttp"
	"html/template"
)

var formTemplate = template.Must(template.New("example").Parse(formTemplateSrc))

func captchaIndex(ctx *clevergo.Context) {
	ctx.SetContentTypeToHTML()
	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	if err := formTemplate.Execute(ctx, &d); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(err.Error())
	}
}

func captchaDisplay(ctx *clevergo.Context) {
	captcha.Server(captcha.StdWidth, captcha.StdHeight).ServeFastHTTP(ctx.RequestCtx)
}

func captchaProcess(ctx *clevergo.Context) {
	ctx.SetContentTypeToHTML()
	if !captcha.VerifyBytes(string(ctx.FormValue("captchaId")), ctx.FormValue("captchaSolution")) {
		fmt.Fprintf(ctx, "Wrong captcha solution! No robots allowed!\n")
	} else {
		fmt.Fprintf(ctx, "Great job, human! You solved the captcha.\n")
	}
	fmt.Fprintf(ctx, "<br><a href='/'>Try another one</a>")
}

func main() {
	app := clevergo.NewApplication()

	// Create a router instance.
	router := app.NewRouter("")

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(captchaIndex))
	router.GET("/captcha/:name", clevergo.HandlerFunc(captchaDisplay))
	router.GET("/process", clevergo.HandlerFunc(captchaProcess))
	router.POST("/process", clevergo.HandlerFunc(captchaProcess))

	// Start server.
	app.Run()
}

const formTemplateSrc = `<!doctype html>
<head><title>Captcha Example</title></head>
<body>
<script>
function setSrcQuery(e, q) {
	var src  = e.src;
	var p = src.indexOf('?');
	if (p >= 0) {
		src = src.substr(0, p);
	}
	e.src = src + "?" + q
}

function playAudio() {
	var le = document.getElementById("lang");
	var lang = le.options[le.selectedIndex].value;
	var e = document.getElementById('audio')
	setSrcQuery(e, "lang=" + lang)
	e.style.display = 'block';
	e.autoplay = 'true';
	return false;
}

function changeLang() {
	var e = document.getElementById('audio')
	if (e.style.display == 'block') {
		playAudio();
	}
}

function reload() {
	setSrcQuery(document.getElementById('image'), "reload=" + (new Date()).getTime());
	setSrcQuery(document.getElementById('audio'), (new Date()).getTime());
	return false;
}
</script>
<select id="lang" onchange="changeLang()">
	<option value="en">English</option>
	<option value="ru">Russian</option>
	<option value="zh">Chinese</option>
</select>
<form action="/process" method=post>
<p>Type the numbers you see in the picture below:</p>
<p><img id=image src="/captcha/{{.CaptchaId}}.png" alt="Captcha image"></p>
<a href="#" onclick="reload()">Reload</a> | <a href="#" onclick="playAudio()">Play Audio</a>
<audio id=audio controls style="display:none" src="/captcha/{{.CaptchaId}}.wav" preload=none>
  You browser doesn't support audio.
  <a href="/captcha/download/{{.CaptchaId}}.wav">Download file</a> to play it in the external player.
</audio>
<input type=hidden name=captchaId value="{{.CaptchaId}}"><br>
<input name=captchaSolution>
<input type=submit value=Submit>
</form>
`
