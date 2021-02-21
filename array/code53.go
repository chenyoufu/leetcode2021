package array

func maxSubArray(nums []int) int {
    sum := nums[0]
    maxSum := nums[0]

    for i := 1; i < len(nums); i++ {
        if sum > 0 {
            sum += nums[i]
        } else {
            sum = nums[i]
        }

        if sum > maxSum {
            maxSum = sum
        }
    }
    return maxSum
}
