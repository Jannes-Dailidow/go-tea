package util

func NewZero[T any]() (v T) { return }

func DerefOrZeroValue[T any](p *T) T {
	if p != nil {
		return *p
	}
	return NewZero[T]()
}
