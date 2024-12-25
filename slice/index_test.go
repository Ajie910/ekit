package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexOf(t *testing.T) {
	testCases := []struct {
		// 测试名称
		name string
		// 切片
		slice []int
		// 查找元素
		val int
		// 期望返回值
		wantVal int
		// 期望的异常返回
		wantErr error
	}{
		{
			name:    "index test 1",
			slice:   []int{123, 100},
			val:     100,
			wantVal: 1,
		},
		{
			name:    "index test 2",
			slice:   []int{123, 124, 125},
			val:     123,
			wantVal: 0,
		},
		{
			name:    "index less than 0",
			slice:   []int{123, 100},
			val:     130,
			wantVal: -1,
		},
		{
			name:    "index last",
			slice:   []int{123, 100, 101, 102, 102, 102},
			val:     102,
			wantVal: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IndexOf(tc.slice, tc.val)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantVal, res)
			fmt.Println(res)
		})
	}
}
