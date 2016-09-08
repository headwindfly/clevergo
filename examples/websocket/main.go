package main

import (
	"github.com/clevergo/websocket"
	"github.com/headwindfly/clevergo"
	"html/template"
	"log"
)

// Create a fasthttp upgrader.
var upgrader = websocket.Upgrader{}

// Create a websocket connection handler.
var connHandler = func(c *websocket.Conn) {
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func echo(ctx *clevergo.Context) {
	upgrader.Upgrade(ctx.RequestCtx, connHandler)
}

func home(ctx *clevergo.Context) {
	ctx.SetContentTypeToHTML()
	homeTemplate.Execute(ctx, "ws://"+string(ctx.Host())+"/echo")
}

func main() {
	app := clevergo.NewApplication()

	// Create a router instance.
	router := app.NewRouter("")

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(home))
	router.GET("/echo", clevergo.HandlerFunc(echo))

	// Start server.
	app.Run()
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server,
"Send" to send a message to the server and "Close" to close the connection.
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
