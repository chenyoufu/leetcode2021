package array

func sortArrayByParity(A []int) []int {

    var i = 0;
    var j = len(A) - 1;
    for ;i<=j;{

        if A[i]%2 > A[j]%2 {
            A[i], A[j] = A[j], A[i]
        }

        if A[i]%2 == 0 {
            i++
        }

        if A[j]%2 != 0 {
            j--
        }
    }

    return A
}

