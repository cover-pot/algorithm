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
