package buffer

import "testing"

func TestBuffer(t *testing.T) {
	cap := 5
	b := NewByteBuffer(cap)
	s := []byte{1, 2}
	b.AppendSlice(s)

	if len(b.data) != len(s) {
		t.Error("buffer does not equal slice")
	}

	t.Logf("buffer: %v\n", b.data)
}
