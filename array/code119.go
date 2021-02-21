package array

func generate(numRows int) [][]int {

    var ret = make([][]int, 0)

    for i := 0; i < numRows; i++ {
        var row = make([]int, i+1)
        row[0] = 1
        row[i] = 1

        for j := 1; j < i; j++ {
            row[j] = ret[i-1][j-1] + ret[i-1][j]
        }

        ret = append(ret, row)
    }

    return ret
}
