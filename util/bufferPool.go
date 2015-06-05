package util

import (
	"bytes"
	"sync"
)

var TextBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 16<<10)) //16KB
	},
}

var ShortTextBufferPool = sync.Pool{
	New: func() interface{}{
		return bytes.NewBuffer(make([]byte,0,8<<10))	//4KB
	},
}