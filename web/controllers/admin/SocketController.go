package admin

import (
	"github.com/kataras/iris/websocket"
	"sync/atomic"
	"fmt"
	"github.com/kataras/iris"
)

type SocketController struct {
	Ctx  iris.Context
	Conn websocket.Connection
}

var visits uint64

func increment() uint64 {
	return atomic.AddUint64(&visits, 1)
}

func (c *SocketController) GetDemo() {
	c.Conn.To(websocket.All).Emit("visit","demo")
	c.Ctx.WriteString("DEMO")
}

func decrement() uint64 {
	return atomic.AddUint64(&visits, ^uint64(0))
}

func (c *SocketController) onLeave(roomName string) {

	fmt.Println("leave")
	//visits--
	newCount := decrement()

	//这将在所有客户端上调用“visit”事件，当前客户端除外
	//（它不能因为它已经离开但是对于任何情况都使用这种类型的设计）
	c.Conn.To(websocket.Broadcast).Emit("visit", newCount)
}

func (c *SocketController) update() {
	// visits++
	newCount := increment()

	//这将在所有客户端上调用“visit”事件，包括当前事件
	//使用'newCount'变量。
	//
	//你有很多方法可以做到更快，例如你可以发送一个新的访问者
	//并且客户端可以自行增加，但这里我们只是“展示”websocket控制器。
	c.Conn.To(websocket.All).Emit("visit", newCount)
}

/* websocket.Connection could be lived here as well, it doesn't matter */
func (c *SocketController) Get() {

	c.Conn.To(websocket.All).EmitMessage([]byte("test"))
	fmt.Println("Get")
	c.Conn.OnLeave(c.onLeave)
	c.Conn.On("visit", c.update)

	// 在所有事件回调注册后调用它。
	c.Conn.Wait()
}

func (c *SocketController) GetTest() {

	fmt.Println("GetTest")
	c.Conn.OnLeave(c.onLeave)
	c.Conn.On("visit", c.update)

	// 在所有事件回调注册后调用它。
	c.Conn.Wait()
}
