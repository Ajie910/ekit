package slice

type Number interface {
	~int | ~int8 | ~int16 | ~int64 | ~float32 | ~float64
}
