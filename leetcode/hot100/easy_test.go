package hot100

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	res := isValid("(}")
	fmt.Println(res)
}
