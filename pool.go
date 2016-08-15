package gcpool

import (
	"log"

	"golang.org/x/net/websocket"
)

type Pool struct {
	conns   map[string]*Conn   // 节点连接
	streams map[string]*Stream // 推送数据
}

func NewPool() *Pool {
	pool := new(Pool)
	pool.conns = make(map[string]*Conn)
	pool.streams = make(map[string]*Stream)
	return pool
}

func (this *Pool) Register(nm string) {
	this.conns[nm] = NewConn(nm)
	this.streams[nm] = NewStream(nm)
}

func (this *Pool) GetConn(nm string) *Conn {
	conn, ok := this.conns[nm]
	if ok {
		return conn
	}
	return nil
}

func (this *Pool) GetStream(nm string) *Stream {
	stream, ok := this.streams[nm]
	if ok {
		return stream
	}
	return nil
}

func (this *Pool) Start() {

	// 为每个流启动一个线程 ...
	for nm, streams := range this.streams {
		go func() {
			stream := streams.Get()
			for {
				select {
				case content := <-stream: // 接收流的数据并推送
					data := content.([2]interface{})
					id := data[0].(string)
					log.Printf("pool stream(%s) content recv(%s)", nm, id)
					this.conns[nm].Get(id, func(conn *websocket.Conn) error {
						log.Printf("pool stream(%s) content send(%s)", nm, id)
						_, err := conn.Write(append(data[1].([]byte), '\n'))
						return err
					})
				}
			}
		}()
	}
}
