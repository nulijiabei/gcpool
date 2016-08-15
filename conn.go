package gcpool

import (
	"golang.org/x/net/websocket"
)

// WebSocket Connect
type Conn struct {
	core *Core
}

// New Connect
func NewConn(name string) *Conn {
	conn := new(Conn)
	conn.core = NewCore(name)
	return conn
}

// Add Connect
func (this *Conn) Add(id string, conn *websocket.Conn) {
	this.core.add(id, conn)
}

// Len Connect
func (this *Conn) Len() int {
	return this.core.len()
}

// Get Connect
func (this *Conn) Get(id string, callback func(conn *websocket.Conn) error) {
	this.core.get(id, func(obj interface{}) error {
		if err := callback(obj.(*websocket.Conn)); err != nil {
			this.core.del(id, func(obj interface{}) {
				obj.(*websocket.Conn).Close()
			})
		}
		return nil
	})
}

// Del Connect
func (this *Conn) Del(id string) {
	this.core.del(id, func(obj interface{}) {
		obj.(*websocket.Conn).Close()
	})
}
