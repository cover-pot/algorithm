package hot100

import (
	"fmt"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	n := 3

	res := generateParenthesis(n)
	fmt.Println(res)
}
