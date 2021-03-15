package array

func isMonotonic(A []int) bool {

    if len(A) == 1 {
        return true
    }

    increasing := true
    for i:=1;i < len(A); i++ {
        if A[i] > A[i-1] {
            increasing = false
            break
        }
    }

    for i:=1;i < len(A); i++ {
        if increasing  {
            if A[i] > A[i-1] {
                return false
            }
        }else {
            if A[i] < A[i-1] {
                return false
            }
        }
    }

    return true

}

func isMonotonic2(A []int) bool {
    increasing := true
    decreasing := true
    for i:=0;i < len(A)-1; i++ {
        if A[i] > A[i+1] {
            increasing = false
        }
        if A[i] < A[i+1] {
            decreasing = false
        }
    }

    return increasing || decreasing
}

