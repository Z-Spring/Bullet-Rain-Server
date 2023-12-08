package msg

import (
	"fmt"
	"testing"
)

func TestDecodeName(t *testing.T) {
	s := make([]byte, 10)
	count := 1
	DecodeName(s, 2, &count)
	fmt.Println(count)
}
