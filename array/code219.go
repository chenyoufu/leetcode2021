package array

func containsNearbyDuplicate(nums []int, k int) bool {

    m := map[int]int{}

    for i, v := range nums {

        _, ok := m[v]
        if ok && i-m[v] <= k {
            m[v] = k
            return true
        } else {
            m[v] = i
        }
    }

    return false

}
