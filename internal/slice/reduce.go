package slice

import "ekit/internal/errors"

// Reduce
// @Description 对切片进行聚合操作
// @Author ajie
// @Param src
// @Param reduce
// @Param init
// @Return T
// @Return error
func Reduce[T, M Number](src []T, reduce func(acc M, cur T) M, initVal M) (M, error) {
	length := len(src)
	if length == 0 {
		var zeroVal M
		return zeroVal, errors.NewErrEmptySlice()
	}
	// 初始值
	acc := initVal
	for _, t := range src {
		acc = reduce(acc, t)
	}
	return acc, nil
}
