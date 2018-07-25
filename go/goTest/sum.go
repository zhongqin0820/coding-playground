package goTest

// add sum of an integer array
func Sum(in []int) (res int) {
	res = 0
	for _, v := range in {
		res += v
	}
	return
}

// return a slice of each incoming slice's sum
func SumAdd(in ...[]int) (res []int) {
	res = make([]int, len(in))
	for i, v := range in {
		res[i] = Sum(v)
	}
	return
}

// return a slice of each incoming slice's sum without the first element
func SumTails(in ...[]int) (Res []int) {
	var res []int
	for _, v := range in {
		if len(v) <= 1 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(v[1:]))
		}
	}
	return res
}
