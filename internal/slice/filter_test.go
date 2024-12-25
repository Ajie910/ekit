package slice

import (
	"ekit/internal/errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		f         func(t int) bool
		wantSlice []int
		wantErr   error
	}{
		{
			name:  "num less than 123",
			slice: []int{123, 100},
			f: func(t int) bool {
				return t < 123
			},
			wantSlice: []int{100},
		},
		{
			name:  "even slice",
			slice: []int{123, 124, 125},
			f: func(t int) bool {
				return t%2 == 0
			},
			wantSlice: []int{124},
		},
		{
			name:  "empty slice",
			slice: []int{},
			f: func(t int) bool {
				return t%2 == 1
			},
			wantErr: errors.NewErrEmptySlice(),
		},
		{
			name:  "num equal 102",
			slice: []int{123, 100, 101, 102, 102, 102},
			f: func(t int) bool {
				return t == 102
			},
			wantSlice: []int{102, 102, 102},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Filter(tc.slice, tc.f)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", res)
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
