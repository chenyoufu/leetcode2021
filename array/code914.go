package array

func hasGroupsSizeX(deck []int) bool {

    if len(deck) < 2 {
        return false
    }

    var m = map[int]int{}

    for _, v := range deck {
        m[v]++
    }

    var x = 0
    for _, v := range m {
        if x == 0 {
            x = v
        } else {
            x = gcd(x, v)
        }
    }

    return x >= 2
}

func gcd(a, b int) int {

    if a == 0 {
        return b
    }

    return gcd(b%a, a)
}
