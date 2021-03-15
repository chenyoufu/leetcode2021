package array

func transpose(matrix [][]int) [][]int {

    m := len(matrix)
    n := len(matrix[0])
    var ret = [][]int{}
    for i:=0;i<n; i++ {
        row := []int{}
        for j:=0;j<m;j++ {
            row = append(row, matrix[j][i])
        }
        ret = append(ret, row)
    }

    return ret
}
