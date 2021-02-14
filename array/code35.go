package array

import "fmt"

func searchInsert(nums []int, target int) int {

    var pos = 0

    for i := 0; i < len(nums); i++ {
        if nums[i] < target {
            pos++
        } else {
            break
        }
    }

    return pos
}

func searchInsert2(nums []int, target int) int {

    var low, high = 0, len(nums) - 1
    var pos = 0

    for low < high {
        mid := low + (high-low)/2
        fmt.Println(low, high, mid)

        if nums[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    if target > nums[low] {
        pos = low + 1
    } else {
        pos = low
    }

    return pos
}
