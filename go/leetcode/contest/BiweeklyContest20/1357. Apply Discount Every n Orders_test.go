package contest

import "testing"

// https://leetcode.com/contest/biweekly-contest-20/problems/apply-discount-every-n-orders/
type Cashier struct {
	n        int
	discount int
	products map[int]int
	cnt      int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	m := make(map[int]int, len(products))
	for i := range products {
		m[products[i]] = prices[i]
	}
	return Cashier{
		n:        n,
		discount: discount,
		products: m,
		cnt:      0,
	}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.cnt++
	var cost float64
	for i, idx := range product {
		cost = cost + float64(this.products[idx]*amount[i])
	}
	if this.cnt%this.n == 0 { // discount
		cost = (1 - float64(this.discount)/100) * cost
	}
	return cost
}

func TestCashier(t *testing.T) {
	// ["Cashier","getBill","getBill","getBill","getBill","getBill","getBill","getBill"]
	// [[3,50,[1,2,3,4,5,6,7],[100,200,300,400,300,200,100]],[[1,2],[1,2]],[[3,7],[10,10]],[[1,2,3,4,5,6,7],[1,1,1,1,1,1,1]],[[4],[10]],[[7,3],[10,10]],[[7,5,3,1,6,4,2],[10,10,10,9,9,9,7]],[[2,3,5],[5,3,2]]]
	obj := Constructor(3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100})
	obj.GetBill([]int{1, 2}, []int{1, 2})
	obj.GetBill([]int{3, 7}, []int{10, 10})
	obj.GetBill([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1, 1, 1, 1, 1})
	obj.GetBill([]int{4}, []int{10})
	obj.GetBill([]int{7, 3}, []int{10, 10})
	obj.GetBill([]int{7, 5, 3, 1, 6, 4, 2}, []int{10, 10, 10, 9, 9, 9, 7})
	obj.GetBill([]int{2, 3, 5}, []int{5, 3, 2})
}

/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
["Cashier","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill","getBill"]
[[192,34,[77],[302]],[[77],[343]],[[77],[990]],[[77],[101]],[[77],[577]],[[77],[160]],[[77],[20]],[[77],[407]],[[77],[205]],[[77],[511]],[[77],[456]],[[77],[287]],[[77],[560]],[[77],[945]],[[77],[453]],[[77],[165]],[[77],[326]],[[77],[222]],[[77],[30]],[[77],[464]],[[77],[916]],[[77],[153]],[[77],[170]],[[77],[368]],[[77],[215]],[[77],[684]],[[77],[21]],[[77],[78]],[[77],[825]],[[77],[259]],[[77],[609]],[[77],[80]],[[77],[660]],[[77],[740]],[[77],[453]],[[77],[918]],[[77],[574]],[[77],[897]],[[77],[135]],[[77],[391]],[[77],[912]],[[77],[560]],[[77],[215]],[[77],[700]],[[77],[557]],[[77],[364]],[[77],[213]],[[77],[331]],[[77],[627]],[[77],[812]],[[77],[84]],[[77],[501]],[[77],[683]],[[77],[603]],[[77],[454]],[[77],[160]],[[77],[19]],[[77],[25]],[[77],[381]],[[77],[595]],[[77],[198]],[[77],[52]],[[77],[734]],[[77],[742]],[[77],[419]],[[77],[555]],[[77],[330]],[[77],[631]],[[77],[132]],[[77],[825]],[[77],[850]],[[77],[923]],[[77],[171]],[[77],[72]],[[77],[13]],[[77],[668]],[[77],[729]],[[77],[64]],[[77],[657]],[[77],[223]],[[77],[981]],[[77],[107]],[[77],[477]],[[77],[111]],[[77],[812]],[[77],[419]],[[77],[391]],[[77],[448]],[[77],[75]],[[77],[842]],[[77],[627]],[[77],[776]],[[77],[297]],[[77],[711]],[[77],[309]],[[77],[654]],[[77],[526]],[[77],[921]],[[77],[73]],[[77],[360]],[[77],[728]],[[77],[499]],[[77],[856]],[[77],[678]],[[77],[488]],[[77],[111]],[[77],[860]],[[77],[352]],[[77],[193]],[[77],[922]],[[77],[859]],[[77],[865]],[[77],[113]],[[77],[370]],[[77],[966]],[[77],[694]],[[77],[432]],[[77],[549]],[[77],[909]]]
*/
