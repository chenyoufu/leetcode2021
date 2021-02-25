package array

func containsDuplicate(nums []int) bool {

    var m = map[int]int{}

    for _, v := range nums {
        m[v]++
        if m[v] > 1 {
            return true
        }
    }

    return false
}
