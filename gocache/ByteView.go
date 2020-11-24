package gocache

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}
func (v ByteView) ByteSlice() []byte {
	return copyByte(v.b)
}
func copyByte(x []byte) []byte {
	c := make([]byte, len(x))
	copy(c, x)
	return c
}
func (v ByteView) String() string {
	return string(v.b)
}
