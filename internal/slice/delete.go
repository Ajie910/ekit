package slice

import "ekit/internal/errors"

// Delete
// @Description 删除指定位置的元素，并返回
// @Author ajie 2024-12-18 11:23:29
// @Param src
// @Param index
// @Return []T
// @Return T
// @Return error
func Delete[T any](src []T, index int) ([]T, T, error) {
	// 获取切片长度
	length := len(src)
	// 校验索引
	if index < 0 || index >= length {
		var zeroVal T
		return nil, zeroVal, errors.NewErrIndexOutOfRange(length, index)
	}
	// 获取要删除的元素
	element := src[index]
	// 初始化目标切片
	ts := make([]T, 0, length-1)
	// 将原始切片前半部分复制到目标切片
	for i := 0; i < index; i++ {
		ts = append(ts, src[i])
	}
	// 将原始切片中要删除元素的后半部分复制到目标切片
	for i := index + 1; i < length; i++ {
		ts = append(ts, src[i])
	}
	return ts, element, nil
}
