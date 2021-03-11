package array

func pivotIndex(nums []int) int {

    var sum = 0
    for _, v := range nums {
        sum += v
    }

    var cur = 0
    for k, v := range nums {

        if cur*2 == sum-v {
            return k
        }

        cur += v
    }

    return -1
}
