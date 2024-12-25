package slice

import "ekit/internal/errors"

// Add [T any] @Title
// @Description  在指定索引处添加元素
// @Author ajie 2024-12-18 11:00:23
// @Param src
// @Param element
// @Param index
// @Return []T
// @Return error
func Add[T any](src []T, element T, index int) ([]T, error) {
	// 获取原切片长度
	length := len(src)
	// 对索引进行判断
	if index < 0 || index > length {
		return nil, errors.NewErrIndexOutOfRange(length, index)
	}

	// 切片末尾进行零值填充
	var zeroVal T
	src = append(src, zeroVal)

	for i := len(src) - 1; i > index; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = element
	return src, nil
}
