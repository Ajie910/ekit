package slice

// Map
// @Description 切片中每个元素进行转换得到新的切片
// @Author ajie
// @Param src
// @Param f
// @Return []M
func Map[M, T any](src []T, f func(T) M) []M {
	ms := make([]M, 0, len(src))
	for _, t := range src {
		ms = append(ms, f(t))
	}

	return ms
}
