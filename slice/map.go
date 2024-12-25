package slice

// Map
// @Description 切片中每个元素进行转换得到新的切片
// @Author ajie
// @Param src
// @Param f
// @Return []M
func Map[Src, Dest any](src []Src, f func(idx int, src Src) Dest) []Dest {
	ms := make([]Dest, 0, len(src))
	for idx, t := range src {
		ms = append(ms, f(idx, t))
	}

	return ms
}
