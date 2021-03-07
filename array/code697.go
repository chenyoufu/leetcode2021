package array

func findShortestSubArray(nums []int) int {

    var m = map[int]int{}
    var startN = map[int]int{}
    var stopN = map[int]int{}
    for k, v := range nums {
        if _, ok := m[v]; ok {
            m[v]++
            stopN[v] = k
        } else {
            m[v] = 1
            startN[v] = k
            stopN[v] = k
        }
    }
    ////
    //fmt.Println(nums)
    //fmt.Println(m, startN, stopN)

    var max = 0
    for _, v := range m {
        if v > max {
            max = v
        }
    }

    //fmt.Println(max)

    var min = len(nums)
    for k, v := range m {
        if v == max {
            if stopN[k]-startN[k] < min {
                min = stopN[k] - startN[k]
            }
        }
    }

    return min+1
}
