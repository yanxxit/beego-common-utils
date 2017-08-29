package comutil

import (
	"testing"
	"fmt"
)

func TestGetPayid(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GetPayid())//长度无法确认
	}
}
