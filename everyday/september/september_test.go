package september

import (
	"fmt"
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
