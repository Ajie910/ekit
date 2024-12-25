package slice

import "ekit/internal/errors"

// Filter
// @Description 过滤切片
// @Author ajie 2024-12-18 11:37:05
// @Param src 原始切片
// @Param f 过滤函数
// @Return []T 目标切片
// @Return error
func Filter[T any](src []T, f func(T) bool) ([]T, error) {
	length := len(src)
	if length == 0 {
		return nil, errors.NewErrEmptySlice()
	}
	ts := make([]T, 0, length)
	for _, t := range src {
		if f(t) {
			ts = append(ts, t)
		}
	}
	return ts, nil
}
