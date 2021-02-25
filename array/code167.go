package array

func twoSum2(numbers []int, target int) []int {
    var m = map[int]int{}

    for k, v := range numbers {
        vv, ok := m[target-v]
        if !ok {
            m[v] = k
        } else {
            return []int{vv+1, k+1}
        }
    }

    return nil
}
