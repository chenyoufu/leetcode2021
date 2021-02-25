package array

func missingNumber(nums []int) int {

    n := len(nums)

    var sum int
    for i:=0;i<n;i++ {
        sum += nums[i]
    }

    return (n +1) *n /2 - sum

}
