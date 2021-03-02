package array

import "sort"

func arrayPairSum(nums []int) int {

    var ret int
    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        if i%2 == 0 {
            ret += nums[i]
        }
    }
    return ret
}
