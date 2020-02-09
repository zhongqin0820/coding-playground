package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func numberOfSteps(num int) int {
	res := 0
	for ; num != 0; res++ {
		switch num % 2 {
		case 0:
			num /= 2
		case 1:
			num -= 1
		}
	}
	return res
}

func TestNumberOfSteps(t *testing.T) {
	tests := []struct {
		num    int
		output int
	}{
		{14, 6},
		{8, 4},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numberOfSteps(ts.num))
		})
	}
}
