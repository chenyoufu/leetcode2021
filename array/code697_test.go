package array

import (
    "fmt"
    "testing"
)

func Test_findShortestSubArray(t *testing.T) {

    nums := []int{1,2}
    s := findShortestSubArray(nums)
    fmt.Println(s)
}