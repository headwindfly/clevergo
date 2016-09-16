# CleverGo
[![Go Report Card](https://goreportcard.com/badge/github.com/headwindfly/clevergo)](https://goreportcard.com/report/github.com/headwindfly/clevergo)
[![GoDoc](https://godoc.org/github.com/headwindfly/clevergo?status.svg)](https://godoc.org/github.com/headwindfly/clevergo)
[![Build Status](https://travis-ci.org/headwindfly/clevergo.svg?branch=master)](https://travis-ci.org/headwindfly/clevergo)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/headwindfly/clevergo.svg?maxAge=2592000)](https://github.com/headwindfly/clevergo/releases)
[![codebeat badge](https://codebeat.co/badges/45b10850-bf4e-40aa-b82a-48d10f2fd5aa)](https://codebeat.co/projects/github-com-headwindfly-clevergo)

**CleverGo** is a **simple**, **high performance** and **secure** web framework for Go (go programing language).
It built on top of [**fasthttp**](https://github.com/valyala/fasthttp).

Some common features and components can be found at [https://github.com/clevergo](https://github.com/clevergo).

Such as: [**websocket**](https://github.com/clevergo/websocket), 
[**sessions**](https://github.com/clevergo/sessions), 
[**captcha**](https://github.com/clevergo/captcha), 
[**csrf**](https://github.com/clevergo/csrf), 
[**jwt**](https://github.com/clevergo/jwt)

1. [**Installation**](#installation)
2. [**Documentation**](#documentation)
3. [**Features**](#features)
4. [**Middlewares**](#middlewares)
5. [**Examples**](#examples)
6. [**Contribution**](#contribution)
7. [**Relevant Packages**](#relevant-packages)


## Installation
```
go get github.com/headwindfly/clevergo
```

### Documentation
**The documentations is not complete now, but we provided completed [examples](#examples).**
- [**English**](docs/en)
- [**中文**](docs/zh)

[Back to top](#readme)


### Features
- **High performance**

1. CleverGo uses **fasthttp** instead of **net/http**, so it is more fast than net/http‘s frameworks.
2. CleverGo's handler [**router**](https://github.com/clevergo/router) is a high performance router(fork from [fasthttprouter](https://github.com/buaazp/fasthttprouter)).
3. Simple architecture.
4. No reflect.

Please refer to [**Go Web Framework Benchmark**](https://github.com/smallnest/go-web-framework-benchmark) for getting more detail.

- **Lightweight**

CleverGo's architecture is very simple, such as the [**Middleware**](middleware.go) and [**Handler**](handler.go).

- **Easy to use**

We provides some examples below, see also [**Examples**](#examples).

- **Multiple Domains**

See also [Application example](examples/application)

[Back to top](#readme)


### Middlewares

| Name                 | Description                                   | Usage                                                                              |
| :---                 | :---------------------------------------------| :----------------------------------------------------------------------------------|
| **Session Middlware**| Session Middleware                            | [**Session Middlware**](middlewares/session)                                       |
| **CSRF Middleware**  | CSRF attack protection                        | [**CSRF Middleware**](middlewares/csrf)                                            |
| **JWT Middleware**   | JSON WEB TOKEN Middleware                     | [**JWT Middleware**](middlewares/jwt)                                              |                       

[Back to top](#readme)

### Examples

| Name                 | Description                                   | Usage                                                                              |
| :---                 | :---------------------------------------------| :----------------------------------------------------------------------------------|
| **Basic Usage**      | Basic Usage                                   | [**Basic Usage**](examples/basic)                                                  |
| **Middleware**       | Middleware                                    | [**Middleware**](examples/middleware)                                              |
| **Websocket**        | Websocket                                     | [**Websocket**](examples/websocket)                                                |
| **Session**          | Session                                       | [**Session**](examples/session)                                                    |
| **RESTFUL API**      | RESTFUL API                                   | [**Restful API**](examples/restful)                                                |
| **CSRF Middleware**  | CSRF attack protection                        | [**CSRF Protection**](examples/csrf)                                               |
| **Captcha**          | Captcha                                       | [**Captcha**](examples/captcha)                                                    |
| **JSON WEB TOKEN**   | JSON WEB TOKEN                                | [**JSON WBE TOKEN**](examples/jwt)                                                 |

More examples can be found at [Examples](examples).

[Back to top](#readme)


### TODO LIST
1. Perfect the documents and examples.
2. Add tests.

[Back to top](#readme)


### Contribution
1. Fork this repository.
2. Added your code on your repository.
3. Send pull request.

**I am willing to accept any pull requests and advises.**

[Back to top](#readme)


### Relevant Packages
Most of packages can be found at https://github.com/clevergo.

- [**fasthttp**](https://github.com/valyala/fasthttp)
- [**router**](https://github.com/clevergo/router)
- [**websocket**](https://github.com/clevergo/websocket)
- [**sessions**](https://github.com/clevergo/sessions)
- [**captcha**](https://github.com/clevergo/captcha)
- [**csrf**](https://github.com/clevergo/csrf)
- [**jwt**](https://github.com/clevergo/jwt)
- [**utils**](https://github.com/clevergo/utils)
- [**pagination**](https://github.com/clevergo/pagination)
- [**i18n**](https://github.com/clevergo/i18n)
- [**assets**](https://github.com/clevergo/assets)

[Back to top](#readme)


### Actual Applications
- [**HeadwindFly.com**](https://github.com/headwindfly/headwindfly.com): https://github.com/headwindfly/headwindfly.com

    1. [Frontend Application](https://headwindfly.com): https://headwindfly.com 
    2. [Backend Application](http://backend.headwindfly.com): http://backend.headwindfly.com 

**How to add my application?**

Fork and added your application in **README.md** and then send pull request.

[Back to top](#readme)
