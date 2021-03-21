package array

func addToArrayForm(A []int, K int) []int {

    var x = []int{}
    for {
        x = append(x, K%10)
        K = K / 10
        if K == 0 {
            break
        }
    }

    var l = len(A)
    if len(x) > len(A) {
        l = len(x)
    }

    var y = make([]int, l)
    for i:=0;i<len(x);i++ {
        y[l-1-i] = x[i]
    }

    var z = make([]int, l)
    for i:=0;i<len(A);i++ {
        z[l-1-i] = A[len(A)-1-i]
    }

    var a []int

    var lastFlag = 0
    for i := len(y) - 1; i >= 0; i-- {
        s := y[i] + z[i]+ lastFlag
        a = append(a, s % 10)
        lastFlag = s / 10
    }

    if lastFlag > 0 {
        a = append(a, lastFlag)
    }

    var res =make([]int, len(a))

    for i:=0;i<len(a);i++ {
        res[len(a)-1-i] = a[i]
    }

    return res
}
