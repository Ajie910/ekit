package slice

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		// 测试名称
		name string
		// 切片1
		s1 []int
		// 切片2
		s2 []int
		// 期望的返回切片
		wantSlice []int
		// 期望的异常返回
		wantErr error
	}{
		{
			name:      "Intersection 1",
			s1:        []int{123, 100, 107, 172, 109, 155, 172},
			s2:        []int{105, 137, 100, 172, 121},
			wantSlice: []int{100, 172},
		},
		{
			name:      "Intersection 2",
			s1:        []int{123, 124, 125},
			s2:        []int{105, 137, 100, 172, 123},
			wantSlice: []int{123},
		},
		{
			name:      "Intersection 3",
			s1:        []int{123, 100, 101, 102, 102, 102},
			s2:        []int{123, 163, 175, 132, 102, 184},
			wantSlice: []int{123, 102},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Intersection(tc.s1, tc.s2)
			fmt.Printf("%+v \n", res)
			//assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		// 测试名称
		name string
		// 切片1
		s1 []int
		// 切片2
		s2 []int
		// 期望的返回切片
		wantSlice []int
		// 期望的异常返回
		wantErr error
	}{
		{
			name:      "Intersection 1",
			s1:        []int{123, 100, 107, 172, 109, 155, 172},
			s2:        []int{105, 137, 100, 172, 121},
			wantSlice: []int{123, 105, 100, 137, 107, 172, 109, 121, 155, 172},
		},
		{
			name:      "Intersection 2",
			s1:        []int{123, 124, 125},
			s2:        []int{105, 137, 100, 172, 123},
			wantSlice: []int{123, 105, 124, 137, 125, 100, 172},
		},
		{
			name:      "Intersection 3",
			s1:        []int{123, 100, 101, 102, 102, 102},
			s2:        []int{123, 163, 175, 132, 102, 184},
			wantSlice: []int{123, 100, 163, 101, 175, 102, 132, 184},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Union(tc.s1, tc.s2)
			fmt.Printf("%+v \n", res)
			//assert.Equal(t, tc.wantSlice, res)
		})
	}
}
