package september

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKWeakestRows(t *testing.T) {

	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	res := KWeakestRows(mat, 3)
	fmt.Println(res)
}

func TestFindDuplicate(t *testing.T) {
	nums := []int{3, 1, 3, 4, 2}

	res := FindDuplicate(nums)
	assert.Equal(t, res, 3)
}

func TestMinOperations(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 4, 2, 3}, 5}, 2},
		{"2", args{[]int{5, 6, 7, 8, 9}, 4}, -1},
		{"3", args{[]int{3, 2, 20, 1, 1, 3}, 10}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MinOperations(tt.args.nums, tt.args.x), "MinOperations(%v, %v)", tt.args.nums, tt.args.x)
		})
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{[]int{1, 3}, []int{2}}, 2.0},
		{"2", args{[]int{1, 2}, []int{3, 4}}, 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FindMedianSortedArrays(tt.args.nums1, tt.args.nums2), "FindMedianSortedArrays(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
