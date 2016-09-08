# CSRF (Cross Site Request Forgery) Protection Middleware
This middleware requires [CleverGo CSRF](https://github.com/clevergo/csrf) and 
[CleverGo Session Middleware](https://github.com/clevergo/middlewares/session).

**This session middleware is necessary.**

## Install dependences
```
go get github.com/clevergo/sessions
go get github.com/clevergo/csrf
```
About the usages of **sessions** package, please refer to https://github.com/clevergo/sessions.

About the usages of **csrf** package, please refer to https://github.com/clevergo/csrf.

## Example
[CSRF Protection Example](/examples/csrf/main.go)
```
go run $GOPATH/src/github.com/headwindfly/clevergo/examples/csrf/main.go
```

## Usages
**Set Error Handler**
```
CSRFMiddleware.SetErrorHandler(handler clevergo.HandlerFunc)
```