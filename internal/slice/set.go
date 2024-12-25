package slice

// Intersection
// @Description 求两个切片的交集
// @Author ajie
// @Param s1
// @Param s2
// @Return []T
func Intersection[T any](s1 []T, s2 []T) []T {
	length1 := len(s1)
	length2 := len(s2)
	var s []T
	if length1 >= length2 {
		s = s1
	} else {
		s = s2
	}
	ts := make([]T, 0, len(s1)+len(s2))
	for _, t1 := range s {
		if idx, err := IndexOf(s2, t1); err == nil && idx >= 0 {
			ts = append(ts, t1)
		}
	}
	return ts
}

// Union
// @Description 取两个切片的并集，并去重
// @Author ajie
// @Param s1
// @Param s2
// @Return []T
func Union[T any](s1 []T, s2 []T) []T {
	// 创建目标切片
	ts := make([]T, 0, len(s1)+len(s2))
	// 求两个切片长度
	length1 := len(s1)
	length2 := len(s2)
	// 记录最长的切片比最短的切片多出来的部分
	var tmpSuffix []T
	var length int
	// 求两个切片的公共长度
	if length1 > length2 {
		tmpSuffix = s1[length2:]
		length = length2
	} else if length1 < length2 {
		tmpSuffix = s2[length1:]
		length = length1
	} else {
		length = length1
	}
	// 遍历公共长度，对两个切片等长的部分进行遍历判断是否加入到目标切片
	var i = 0
	for ; i < length; i++ {
		if idx, err := IndexOf(ts, s1[i]); err == nil && idx < 0 {
			ts = append(ts, s1[i])
		}
		if idx, err := IndexOf(ts, s2[i]); err == nil && idx < 0 {
			ts = append(ts, s2[i])
		}
	}
	// 如果存在超出切片，对超出的切片继续判断
	if tmpSuffix != nil {
		for i := 0; i < len(tmpSuffix); i++ {
			if idx, err := IndexOf(ts, tmpSuffix[i]); err == nil && idx < 0 {
				ts = append(ts, tmpSuffix[i])
			}
		}
	}
	return ts
}

// Subtract
// @Description 求两个切片的差集，第一个切片减去第二个
// @Author ajie
// @Param s1
// @Param s2
// @Return []T
func Subtract[T any](s1 []T, s2 []T) []T {
	var ts []T
	for _, t := range s2 {
		idx, err := IndexOf(s1, t)
		if err == nil && idx >= 0 {
			ts, _, _ = Delete(s1, idx)
			if len(s1) == 0 {
				return ts
			}
		}
	}
	return ts
}
