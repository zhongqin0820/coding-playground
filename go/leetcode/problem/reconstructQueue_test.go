package problemSet

import (
    "log"
    "testing"
)


func TestReconstructQueue(t *testing.T) {
    t.Run("reconstructQueue", func(t *testing.T){
        p := [][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}
        res := reconstructQueue(p)
        log.Println(res)
    })
}
