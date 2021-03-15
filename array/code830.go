package array

func largeGroupPositions(s string) [][]int {

    last := s[0]
    start := 0
    end := 0

    var ret = [][]int{}
    i := 0

    for {

        if i == len(s) || s[i] != last {
            end = i - 1
            if i-start >= 3 {
                ret = append(ret, []int{start, end})
            }

            if i == len(s) {
                break
            }
            start = i
            last = s[i]
        }

        i++
    }

    return ret
}
