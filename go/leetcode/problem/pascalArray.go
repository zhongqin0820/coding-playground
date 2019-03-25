package problemSet

import (
	"log"
)

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}
	var res [][]int
	for i := 0; i < numRows; i++ {
		var subRes []int = make([]int, i+1, i+1)
		for j := 0; j <= i; j++ {
			log.Println(i, j)
			if j == 0 || j == i {
				subRes[j] = 1
			} else {
				subRes[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res = append(res, subRes)
		log.Println(res)
	}
	return res
}

// func main3() {
// 	Res := generate(5)
// 	log.Println(Res)
// }
