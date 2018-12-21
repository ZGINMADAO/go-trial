package socket

import (
	"github.com/kataras/iris/websocket"
	"github.com/kataras/iris"
	"fmt"
)

func StartSocket(app *iris.Application) {
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})

	ws.OnConnection(handleConnection)

	app.Get("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(ws.ClientSource)
	})

	app.Get("/echo", ws.Handler())
}

func handleConnection(c websocket.Connection) {
	// Read events from browser
	c.On("chat", func(msg string) {
		// Print the message to the console, c.Context() is the iris's http context.
		fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
		// Write message back to the client message owner with:
		// c.Emit("chat", msg)
		// Write message to all except this client with:
		c.To(websocket.Broadcast).Emit("chat", msg)
	})

	c.OnMessage(func(bytes []byte) {
		c.To(c.ID()).EmitMessage(bytes)
		fmt.Println(string(bytes))
	})
}
