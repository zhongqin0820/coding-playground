// Q1
func legal(a, b int) bool {
    if a < 1 || a > 100000 {
        return false
    }
    if b < 1 || b > 1000000000 {
        return false
    }
    return true
}

func main() {
    a := 0
    b := 0
    // for {
    fmt.Scan(&a, &b)
    if legal(a, b) == false {
        return
    }
    fmt.Println(b/a + 1)
    // }
}


//Q2
func judge(a, b int) bool {
    if a < 1 || a > 1000000000 || a > b {
        return false
    }
    if b < 1 || b > 1000000000 {
        return false
    }
    return true
}

func judgeP(n int) bool {
    if n < 1 || n > 100000 {
        return false
    }
    return true
}

func calRes(a, b int) int {
    // 相等,偶数
    if a == b && a%2 == 0 {
        return a
    } else if a == b && a%2 == 1 { //相等,奇数
        return -a
    }
    if a%2 == b%2 && a%2 == 0 {
        return b - (b-a)/2
    } else if a%2 == b%2 && a%2 == 1 {
        return -b + (b-a)/2
    } else if b%2 == 0 && a%2 == 1 {
        return (b-a)/2 + 1
    }
    return -(b-a)/2 - 1
}

func main() {
    n := 0
    a, b := 0, 0
    fmt.Scan(&n)
    if judgeP(n) == false {
        return
    }
    for ; n > 0; n-- {
        fmt.Scan(&a, &b)
        if judge(a, b) == false {
            return
        }
        fmt.Println(calRes(a, b))
    }
}

//Q3
func judge(a, b int) bool {
    if a < 1 || a > 2000 {
        return false
    }
    if b < 0 || b > 2000 {
        return false
    }
    return true
}

func judgeNum(n int) bool {
    if n < 0 || n > 2 {
        return false
    }
    return true
}

func sort(s []int, flag bool) []int {
    if len(s) <= 1 {
        return s
    }
    var i, j, v, k = 0, len(s) - 1, s[0], 1

    for i < j {
        if flag { //big->small
            if s[k] < v {
                s[k], s[j] = s[j], s[k]
                j--
            } else {
                s[i], s[k] = s[k], s[i]
                i++
                k++
            }
        } else { //small -> big
            if s[k] > v {
                s[k], s[j] = s[j], s[k]
                j--
            } else {
                s[i], s[k] = s[k], s[i]
                i++
                k++
            }
        }
    }
    sort(s[:i], flag)
    sort(s[i+1:], flag)
    return s
}

func main() {
    n, s := 0, 0
    fmt.Scan(&n, &s)
    if judge(n, s) == false {
        return
    }
    a := []int{}
    var temp int
    for ; n > 0; n-- {
        fmt.Scan(&temp)
        if judgeNum(temp) == false {
            return
        }
        a = append(a, temp)
    }
    a = sort(a, false)
    res := 1
    for _, v := range a {
        if s == 0 {
            break
        }
        if v == 0 {
            res = res * 3
        } else if v == 1 {
            res = res * 2
        } else {
            res = res * 2
        }
        s--
    }
    fmt.Println(res % 10e7)
}
//
