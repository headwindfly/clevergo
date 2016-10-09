# CleverGo
[![Go Report Card](https://goreportcard.com/badge/github.com/headwindfly/clevergo)](https://goreportcard.com/report/github.com/headwindfly/clevergo)
[![GoDoc](https://godoc.org/github.com/headwindfly/clevergo?status.svg)](https://godoc.org/github.com/headwindfly/clevergo)
[![Build Status](https://travis-ci.org/headwindfly/clevergo.svg?branch=master)](https://travis-ci.org/headwindfly/clevergo)
[![Coverage Status](https://coveralls.io/repos/github/headwindfly/clevergo/badge.svg?branch=master)](https://coveralls.io/github/headwindfly/clevergo?branch=master)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/headwindfly/clevergo.svg?maxAge=2592000)](https://github.com/headwindfly/clevergo/releases)
[![codebeat badge](https://codebeat.co/badges/45b10850-bf4e-40aa-b82a-48d10f2fd5aa)](https://codebeat.co/projects/github-com-headwindfly-clevergo)

:point_right: [English](README.md)

**CleverGo** - Go语言编写的简单，高新能，安全的WEB框架，其非常适于设计RESTful API。

CleverGo 不提供**ORM**和**模板引擎**，一是因为不想重复造轮子，二是可以自由选择喜欢的**ORM**和**模板引擎**。

**如果对这个项目有兴趣，非常欢迎一起来维护:smile:。**

另外：英文大多是机翻的，所以英文说明可能不通顺。。。欢迎纠错:smile:。

## 特性
- **高性能**

1. 基于 [**fasthttp**](https://github.com/valyala/fasthttp)
2. [**高性能路由器**](https://github.com/clevergo/router)
3. 不使用Reflect

- **简单**

CleverGo's的结构很简单, 比如 [**Middleware**](clevergo.go) 和 [**Handler**](clevergo.go).

- **易用**


## 性能

![Benchmark](https://github.com/smallnest/go-web-framework-benchmark/blob/master/benchmark.png)

点击右侧链接可获得更多的性能测试信息： [**Go Web Framework Benchmark**](https://github.com/smallnest/go-web-framework-benchmark).


## 安装
```
go get github.com/headwindfly/clevergo
```


## 文档
文档还不完整，**但是已经提供了几乎完整的[例子](https://github.com/clevergo/examples).**
- [**English**](docs/en)
- [**中文**](docs/zh)


## Middlewares

| Name                 | Description                                   | Usage                                                                              |
| :---                 | :---------------------------------------------| :----------------------------------------------------------------------------------|
| **Session Middlware**| Session Middleware                            | [**Session Middlware**](https://github.com/clevergo/sessionmiddleware)             |
| **CSRF Middleware**  | CSRF attack protection                        | [**CSRF Middleware**](https://github.com/clevergo/csrfmiddleware)                  |
| **JWT Middleware**   | JSON WEB TOKEN Middleware                     | [**JWT Middleware**](https://github.com/clevergo/jwtmiddleware)                    |                       

目前中间件比较少，只提供一些常用的中间件，日后会逐步添加和完善。

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

点击右侧链接查看更多例子： [Examples](https://github.com/clevergo/examples).


## Contribution
1. Fork 当前的仓库
2. 在fork的仓库修改代码
3. 发送 pull request.
4. 等待合并
