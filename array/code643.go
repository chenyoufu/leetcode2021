package array

import "math"

func findMaxAverage(nums []int, k int) float64 {

    var max = math.MinInt64
    for i := 0; i < len(nums)-k+1; i++ {
        var sum int
        for j := 0; j < k; j++ {
            sum += nums[i+j]
        }

        if max < sum {
            max = sum
        }
    }

    return float64(max) / float64(k)
}


func findMaxAverage2(nums []int, k int) float64 {

    var max = math.MinInt64

    var sum int
    for i := 0; i < k; i++ {
        sum += nums[i]
    }

    max = sum

    for i := k; i < len(nums); i++ {

        sum += nums[i] - nums[i-k]

        if sum > max {
            max = sum
        }
    }

    return float64(max) / float64(k)
}
