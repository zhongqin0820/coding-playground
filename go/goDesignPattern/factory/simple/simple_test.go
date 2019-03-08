package simple

import (
	"log"
	"testing"
)

func Test_Payments(t *testing.T) {
	for _, idx := range []int{1, 2, 3} {
		p, err := GetPayment(idx)
		if err != nil {
			log.Println(err)
		} else {
			r := p.Pay(123.0)
			log.Println(r)
		}
	}
}
