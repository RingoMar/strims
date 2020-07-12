package pool

// DefaultPool ...
var DefaultPool = New(8)

// Get ...
func Get(size uint16) (b []byte) {
	return DefaultPool.Get(size)
}

// Put ...
func Put(b []byte) {
	DefaultPool.Put(b)
}