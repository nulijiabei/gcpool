package gcpool

import (
	"log"
)

type Stream struct {
	stream chan interface{}
	name   string
}

func NewStream(name string) *Stream {
	stream := new(Stream)
	stream.name = name
	stream.stream = make(chan interface{}, 1000)
	return stream
}

func (this *Stream) Add(id string, data interface{}) {
	log.Printf("Stream(%s) add(%s)", this.name, id)
	this.stream <- [2]interface{}{id, data}
}

func (this *Stream) Get() chan interface{} {
	return this.stream
}
