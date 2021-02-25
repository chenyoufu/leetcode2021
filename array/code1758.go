package array

import "fmt"

func minOperations(s string) int {
    if len(s) == 1 {
        return 0
    }

    var needChange bool
    var count int

    for i := 1; i < len(s); i++ {

        prev := s[i-1]
        if needChange {
            if prev == '0' {
                prev = '1'
            } else {
                prev = '0'
            }
        }

        if s[i] == prev  {
            needChange = true
            count++
        } else {
            needChange = false
        }

        fmt.Println("i:", i, "s[i]:", s[i]," needChange:", needChange, " count:", count, "prev", prev)
    }
    return count
}
