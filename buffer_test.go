package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuffer(t *testing.T) {
	assert := assert.New(t)

	cap := 5
	b := NewByteBuffer(cap)
	s := []byte{50, 51, 52, 53, 54}

	b.AppendSlice(s)

	assert.Equal(len(s), len(b.data))
	assert.True(b.IsFull())
	assert.False(b.IsEmpty())

	v := b.PopRight()

	assert.Equal(s[4], v)
	assert.Len(b.data, cap-1)

	v2 := b.PopLeft()

	assert.Equal(s[0], v2)
	assert.Len(b.data, cap-2)
}
