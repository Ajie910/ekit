package ekit

// ToPtr
// @Description  获取任意类型的指针
// @Author ajie
// @Param t
// @Return *T
func ToPtr[T any](t T) *T {
	return &t
}
