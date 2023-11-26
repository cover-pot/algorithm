package token

import (
	"fmt"
	"testing"
	"time"
)

func TestToken_Generate(t1 *testing.T) {
	token := NewToken()
	for i := 0; i < 200; i++ {
		go func() {
			err := token.Sub()
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	time.Sleep(3 * time.Second)
}
