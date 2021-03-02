package array

import "math"

func thirdMax(nums []int) int {
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }

    maxCount := 0
    for _, v := range nums {
        if v == max {
            maxCount++
        }
    }

    secondMax := nums[0]
    for _, v := range nums {
        if v != max {
            secondMax = v
            break
        }
    }

    for _, v := range nums {
        if v != max && v > secondMax {
            secondMax = v
        }
    }

    secondCount := 0
    for _, v := range nums {
        if v == secondMax {
            secondCount++
        }
    }

    if secondCount + maxCount == len(nums) {
        return max
    }

    thirdMax := nums[0]
    for _, v := range nums {
        if v != max && v!= secondMax{
            thirdMax = v
            break
        }
    }
    for _, v := range nums {
        if v != max && v != secondMax && v > thirdMax {
            thirdMax = v
        }
    }

    return thirdMax
}

func thirdMax2(nums []int) int {
    max := math.MinInt64
    second := math.MinInt64
    third := math.MinInt64

    for _, v := range nums {
        if v > max {
            third = second
            second = max
            max = v
        } else if v>second && v<max {
            third = second
            second = v
        } else if v > third && v<second {
            third = v
        }
    }

    if third == math.MinInt64 {
        return max
    }

    return third
}
