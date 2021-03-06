package array

import "math"

func maximumProduct(nums []int) int {

    if len(nums) == 3 {
        return nums[0] * nums[1] * nums[2]
    }

    var max = math.MinInt64
    var second = math.MinInt64
    var third = math.MinInt64

    var min = math.MaxInt64
    var minSecond = math.MinInt64

    for i := 0; i < len(nums); i++ {
        if nums[i] > max {
            third = second
            second = max
            max = nums[i]
        } else if nums[i] > second && nums[i] <= max {
            third = second
            second = nums[i]
        } else if nums[i] > third && nums[i] <= second {
            third = nums[i]
        }

        if nums[i] < min {
            minSecond = min
            min = nums[i]
        } else if nums[i] >= min && nums[i] < minSecond {
            minSecond = nums[i]
        }
    }

    if max < 0 {
        return max * second * third
    }

    v1 := max * second * third
    v2 := max * min * minSecond

    if v1 > v2 {
        return v1
    }

    return v2
}
