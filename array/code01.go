package array

func twoSum(nums []int, target int) []int {

    m := make(map[int]int, len(nums))
    var ret []int

    for k, v := range nums {
        vv, ok := m[target-v]
        if ok {
            ret = []int{vv, k}
            break
        }
        m[v] = k
    }

    return ret
}
