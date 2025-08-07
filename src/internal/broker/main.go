package broker

import (
	"sync"
)

var (
	_once sync.Once
)

type IBroker interface{}

func New() IBroker {

	var b IBroker
	_once.Do(func() {
		b := &broker{}
		b.init()
	})

	return b
}
