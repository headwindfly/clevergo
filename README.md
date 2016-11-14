# CleverGo

### It is recommend to use [Gem](https://github.com/go-gem/gem) framework.

[![Go Report Card](https://goreportcard.com/badge/github.com/headwindfly/clevergo)](https://goreportcard.com/report/github.com/headwindfly/clevergo)
[![GoDoc](https://godoc.org/github.com/headwindfly/clevergo?status.svg)](https://godoc.org/github.com/headwindfly/clevergo)
[![Build Status](https://travis-ci.org/headwindfly/clevergo.svg?branch=master)](https://travis-ci.org/headwindfly/clevergo)
[![Coverage Status](https://coveralls.io/repos/github/headwindfly/clevergo/badge.svg?branch=master)](https://coveralls.io/github/headwindfly/clevergo?branch=master)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/headwindfly/clevergo.svg?maxAge=2592000)](https://github.com/headwindfly/clevergo/releases)
[![codebeat badge](https://codebeat.co/badges/45b10850-bf4e-40aa-b82a-48d10f2fd5aa)](https://codebeat.co/projects/github-com-headwindfly-clevergo)

:point_right: [中文介绍](README_ZH.md)

**CleverGo** is a **simple**, **high performance** and **secure** web framework for Go (golang programing language).
It built on top of [**fasthttp**](https://github.com/valyala/fasthttp).


1. [**Features**](#features)
2. [**Performance**](#performance)
3. [**Installation**](#installation)
4. [**Documentation**](#documentation)
5. [**Middlewares**](#middlewares)
6. [**Examples**](#examples)
7. [**Contribution**](#contribution)
8. [**Actual Applications**](#actual-applications)


## Features
- **High performance**

1. CleverGo uses [**fasthttp**](https://github.com/valyala/fasthttp) instead of **net/http**, so it is more fast than net/http‘s frameworks.
2. CleverGo's [**router**](https://github.com/clevergo/router) - a high performance router.
3. Simple architecture.
4. No reflect.

Please refer to [**Go Web Framework Benchmark**](https://github.com/smallnest/go-web-framework-benchmark) for getting more detail.

- **Simple**

CleverGo's architecture is very simple, such as the [**Middleware**](clevergo.go) and [**Handler**](clevergo.go).

- **Easy to use**

We provides some examples below, see also [**Examples**](#examples).

[Back to top](#readme)


## Performance

![Benchmark](https://github.com/smallnest/go-web-framework-benchmark/blob/master/benchmark.png)

Further information is available in [**Go Web Framework Benchmark**](https://github.com/smallnest/go-web-framework-benchmark).

[Back to top](#readme)

## Installation
```
go get github.com/headwindfly/clevergo
```


## Documentation
**The documentations is not complete now, but we provided completed [examples](#examples).**
- [**English**](docs/en)
- [**中文**](docs/zh)

[Back to top](#readme)


## Middlewares

| Name                 | Description                                   | Usage                                                                              |
| :---                 | :---------------------------------------------| :----------------------------------------------------------------------------------|
| **Session Middlware**| Session Middleware                            | [**Session Middlware**](https://github.com/clevergo/sessionmiddleware)             |
| **CSRF Middleware**  | CSRF attack protection                        | [**CSRF Middleware**](https://github.com/clevergo/csrfmiddleware)                  |
| **JWT Middleware**   | JSON WEB TOKEN Middleware                     | [**JWT Middleware**](https://github.com/clevergo/jwtmiddleware)                    |                       

[Back to top](#readme)


## Examples

| Name                 | Description                                   | Usage                                                                              |
| :---                 | :---------------------------------------------| :----------------------------------------------------------------------------------|
| **Basic Usage**      | Basic Usage                                   | [**Basic Usage**](https://github.com/clevergo/examples/tree/master/basic)          |
| **Application**      | Application                                   | [**Application**](https://github.com/clevergo/examples/tree/master/application)    |
| **Middleware**       | Middleware                                    | [**Middleware**](https://github.com/clevergo/examples/tree/master/middleware)      |
| **Websocket**        | Websocket                                     | [**Websocket**](https://github.com/clevergo/examples/tree/master/websocket)        |
| **Session**          | Session                                       | [**Session**](https://github.com/clevergo/examples/tree/master/session)            |
| **RESTFUL API**      | RESTFUL API                                   | [**Restful API**](https://github.com/clevergo/examples/tree/master/restful)        |
| **CSRF Middleware**  | CSRF attack protection                        | [**CSRF Protection**](https://github.com/clevergo/examples/tree/master/csrf)       |
| **Captcha**          | Captcha                                       | [**Captcha**](https://github.com/clevergo/examples/tree/master/captcha)            |
| **JSON WEB TOKEN**   | JSON WEB TOKEN                                | [**JSON WBE TOKEN**](https://github.com/clevergo/examples/tree/master/jwt)         |

More examples can be found at [Examples](https://github.com/clevergo/examples).

[Back to top](#readme)


## Contribution
1. Fork this repository.
2. Added your code on your repository.
3. Send pull request.

[Back to top](#readme)


## Relevant Packages
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


## Actual Applications
- [**HeadwindFly.com**](https://github.com/headwindfly/headwindfly.com): https://github.com/headwindfly/headwindfly.com

    1. [https://headwindfly.com](http://headwindfly.com)
    2. [https://docs.headwindfly.com](https://docs.headwindfly.com)
    3. [https://accounts.headwindfly.com](https://accounts.headwindfly.com)
    4. [https://backend.headwindfly.com](https://backend.headwindfly.com)
    5. [https://helpers.headwindfly.com](https://helpers.headwindfly.com)
    6. [https://api.headwindfly.com](https://api.headwindfly.com)

**How to add my application?**

Fork and added your application in **README.md** and then send pull request.

[Back to top](#readme)
