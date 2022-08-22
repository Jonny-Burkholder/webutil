package buf

import (
	"bytes"
	"sync"
)

// bufPool is a buffer pool (for use with templates)
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// GetBuf is just a helper method to make life easier
func GetBuf() *bytes.Buffer {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	return b
}

// PutBuf is just a helper to make life easier
func PutBuf(b *bytes.Buffer) {
	bufPool.Put(b)
}
