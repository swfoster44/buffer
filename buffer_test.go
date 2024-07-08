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

	err := b.AppendSlice(s)
	assert.Nil(err)
	err1 := b.AppendSlice([]byte{55})
	assert.Error(err1)

	v3, err4 := b.PeekLeft(1)
	assert.Equal(s[3], v3[0])
	assert.Nil(err4)

	v4, err5 := b.PeekRight(1)
	assert.Equal(s[1:2], v4)
	assert.Nil(err5)

	assert.False(b.HasSpace(1))

	assert.Equal(len(s), len(b.data))
	assert.True(b.IsFull())
	assert.False(b.IsEmpty())

	v, err2 := b.PopRight()
	assert.Equal(s[4], v)
	assert.Equal(b.Cap(), cap)
	assert.Nil(err2)

	v2, err3 := b.PopLeft()
	assert.Equal(s[0], v2)
	assert.Equal(b.Cap(), cap)
	assert.Nil(err3)

	assert.True(b.HasSpace(2))
}
