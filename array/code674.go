package array

func findLengthOfLCIS(nums []int) int {

    if len(nums) == 0 {
        return 0
    }

    var max = 1
    var cur = 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            cur++
        } else {
            cur = 1
        }

        if cur > max {
            max = cur
        }
    }

    return max
}
