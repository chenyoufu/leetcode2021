package array

func sumEvenAfterQueries(A []int, queries [][]int) []int {

    var res []int
    for i:=0;i<len(queries);i++ {
        val := queries[i][0]
        index := queries[i][1]

        A[index] +=  val

        sum := 0
        for _, v := range A {
            if v % 2 == 0 {
                sum += v
            }
        }
        res = append(res, sum)
    }

    return res
}

func sumEvenAfterQueries2(A []int, queries [][]int) []int {

    var res []int
    var sum = 0

    for _, v := range A {
        if v % 2 == 0 {
            sum += v
        }
    }

    for i:=0;i<len(queries);i++ {
        val := queries[i][0]
        index := queries[i][1]

        if A[index] % 2 == 0 {
            sum -= A[index]
        }

        A[index] +=  val

        if A[index] % 2 == 0 {
            sum += A[index]
        }

        res = append(res, sum)
    }

    return res
}