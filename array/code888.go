package array

func fairCandySwap(A []int, B []int) []int {

    sumA := 0
    for _, v := range A {
        sumA += v
    }

    sumB := 0
    for _, v := range B {
        sumB += v
    }

    // sumA + B - A = sumB - B + A
    // sumA-sumB = 2A-2B
    // 2B = 2A - sumA + sumB

    var m = map[int]int{}

    for _, v := range A {
        m[2*v-sumA+sumB] = v
    }

    for _, v := range B {
        if vv, ok := m[2*v]; ok {
            return []int{vv, v}
        }
    }

    return nil
}
