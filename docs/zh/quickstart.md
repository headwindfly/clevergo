# 快速开始 Quick Start

## 安装
```
go get github.com/headwindfly/clevergo
```

## Hello World

```
package main

import (
	"log"
	"github.com/headwindfly/clevergo"
)


func helloCleverGo(ctx *clevergo.Context) {
	ctx.SetBodyString("Hello CleverGo.")
}

func main() {
    // 创建 Application
    app := clevergo.NewApplication()
    
	// 创建 Router
	router := app.NewRouter("")

	// 注册路由处理器
	router.GET("/", clevergo.HandlerFunc(helloCleverGo))

	// 启动 Server
	app.Run()
}
```

然后访问 http://127.0.0.1:8080 即可看到“Hello CleverGo.”字样。

## 使用案例
[Examples](/examples/basic)

## Shortcut
* [目录](README.md)