package array

func matrixReshape(nums [][]int, r int, c int) [][]int {

    if len(nums) == 0 || len(nums[0]) == 0 {
        return nums
    }

    row := len(nums)
    col := len(nums[0])

    if row*col != r*c {
        return nums
    }

    var temp []int

    for _, x := range nums {
        temp = append(temp, x...)
    }

    var ret [][]int

    for i := 0; i < len(temp); i += c {
        ret = append(ret, temp[i:i+c])
    }

    return ret
}
