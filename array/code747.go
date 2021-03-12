package array

func dominantIndex(nums []int) int {

    var max int
    var index int
    for k, v := range nums {
        if v > max {
            max = v
            index = k
        }
    }

    for _, v := range nums {
        if max == v {
            continue
        }

        if v *2 > max {
            return -1
        }
    }

    return index
}
