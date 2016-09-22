# 验证码 Captcha
防止机器人模拟请求，验证码是个不错的解决方案。所以CleverGo也提供了Captcha的支持。

## 实例
```
go run $GOPATH/src/github.com/headwindfly/clevergo/examples/captcha/main.go
```
[Captcha Example](/examples/captcha)

## 说明
目前的实例只是简单的例子，一般来说:

1. 验证码需要存储到Session里面，验证成功则刷新验证码

2. 允许客户端刷新验证码

目前这个例子，并没有使用Session来存储，以后有空会将其改进，如果能贡献你的代码就最好不过了:)

## Shortcut
* [目录](README.md)
* [Captcha包](https://github.com/clevergo/captcha)
* [Session](session.md)