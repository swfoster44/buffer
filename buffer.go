package buffer

const newLen = 0
const emptyLen = 0
const emptyCap = 0

type Buffered[T any] interface {
	Len() int
	Cap() int
	IsFull() bool
	IsEmpty() bool
	AppendSlice(s []T)
	Append(s T)
	PopLeft() T
	PopRight() T
	PeekLeft() []T
	PeekRight() []T
	InBounds(i int)
	Data(d []T)
	HasSpace(l int)
	Copy() []T
	DataSlice() []T
}

func NewByteBuffer(cap int) *Buffer[byte] {
	d := make([]byte, newLen, cap)
	b := Buffer[byte]{data: d}
	return &b
}

type Buffer[T any] struct {
	data []T
}

func (b *Buffer[T]) Data(data []T) {
	b.data = data
}

func (b *Buffer[T]) DataSlice(i, j int) []T {
	if b.InBounds(i) && b.InBounds(j) {

		return b.data[i:j]
	}

	panic("out of bounds range")
}

func (b *Buffer[T]) HasSpace(l int) bool {
	nl := b.Len() + l
	return nl <= b.Cap()
}

func (b *Buffer[T]) Copy() []T {
	d := make([]T, b.Len(), b.Cap())
	copy(d, b.data)
	return d
}

func (b *Buffer[T]) AppendSlice(s []T) error {
	if b.HasSpace(len(s)) {
		b.data = append(b.data, s...)
		return nil
	}

	return newCapError("AppendSlice()")

}

func (b *Buffer[T]) Append(s T) error {
	if b.HasSpace(1) {
		b.data = append(b.data, s)
		return nil
	}
	return newCapError("Append()")

}

func (b *Buffer[T]) PopRight() (T, error) {
	if !b.IsEmpty() {
		i := 0
		j := b.Len() - 1
		v := b.data[j]

		d := b.data[i:j]
		nd := make([]T, newLen, b.Cap())
		b.data = append(nd, d...)
		return v, nil
	}

	t := new(T)
	return *t, newBuffEmptyError("PopRight()")
}

func (b *Buffer[T]) PopLeft() (T, error) {
	if !b.IsEmpty() {
		i := 0
		v := b.data[i]
		d := b.data[i+1:]
		nd := make([]T, newLen, b.Cap())
		b.data = append(nd, d...)
		return v, nil
	}
	t := new(T)
	return *t, newBuffEmptyError("PopLeft()")

}

func (b *Buffer[T]) LastIndex() int {
	return b.Len() - 1
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
	return b.Len() == emptyLen
}

func (b *Buffer[T]) InBounds(i int) bool {
	return i >= newLen && i <= b.Len()
}

func (b *Buffer[T]) PeekLeft(n int) ([]T, error) {
	j := b.Len() - 1
	i := j - n
	if b.InBounds(i) {
		return b.data[i:j], nil
	}
	t := make([]T, emptyLen, emptyCap)
	return t, newBuffBoundsErr("PeekLeft()")
}

func (b *Buffer[T]) PeekRight(n int) ([]T, error) {
	i := 0
	j := i + n
	if b.InBounds(j) {
		return b.data[i:j], nil
	}
	t := make([]T, emptyLen, emptyCap)
	return t, newBuffBoundsErr("PeekRight()")

}
