package buffer

type Buffered[T any] interface {
	Data() []T
	Len() int
	Cap() int
	IsFull() bool
	IsEmpty() bool
	AppendSlice(s []T)
	Append(s T)
}

type Buffer[T any] struct {
	data []T
}

func (b *Buffer[T]) Data() []T {
	d := make([]T, b.Len(), b.Cap())
	copy(d, b.data)
	return d
}

func (b *Buffer[T]) Len() int {
	return len(b.data)
}

func (b *Buffer[T]) Cap() int {
	return cap(b.data)
}

func (b *Buffer[T]) IsFull() bool {
	return b.Cap() == b.Len()
}

func (b *Buffer[T]) IsEmpty() bool {
	return b.Len() == 0
}

func (b *Buffer[T]) AppendSlice(s []T) {
	if !b.IsFull() {
		b.data = append(b.data, s...)
	} else {

		panic("AppendSlice(s []T): buffer is full")
	}
}

func (b *Buffer[T]) Append(s T) {
	if !b.IsFull() {
		b.data = append(b.data, s)
	} else {
		panic("Append(s T): buffer is full")
	}
}

func NewByteBuffer(cap int) Buffer[byte] {
	d := make([]byte, 0, cap)
	return Buffer[byte]{data: d}
}
