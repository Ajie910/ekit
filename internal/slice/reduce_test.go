package slice

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestReduce(t *testing.T) {
	testCases := []struct {
		// 测试名称
		name string
		// 切片
		slice []int
		// 聚合函数
		f func(prev, cur int) int
		// 初始值
		init int
		// 聚合结果
		val int
		// 期望的异常返回
		wantErr error
	}{
		{
			name:  "reduce sum val 2",
			slice: []int{123, 100},
			f: func(prev int, cur int) int {
				return prev + cur
			},
			init: 0,
			val:  223,
		},
		{
			name:  "reduce sum val 3",
			slice: []int{123, 124, 125},
			f: func(prev int, cur int) int {
				return prev + cur
			},
			init: 0,
			val:  372,
		},
		{
			name:  "reduce sum val with init",
			slice: []int{123, 100, 101, 102, 102, 102},
			f: func(prev int, cur int) int {
				return prev + cur
			},
			init: 22,
			val:  652,
		},
		{
			name:  "reduce max val",
			slice: []int{123, 100, 101, 102, 102, 102},
			f: func(prev int, cur int) int {
				return int(math.Max(float64(prev), float64(cur)))
			},
			init: 0,
			val:  123,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Reduce(tc.slice, tc.f, tc.init)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			println(res)
			assert.Equal(t, tc.val, res)
		})
	}
}
