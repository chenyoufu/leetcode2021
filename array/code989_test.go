package array

import (
    "fmt"
    "testing"
)

func Test_addToArrayForm(t *testing.T) {
    A := []int{215}
    K := 806
    ret := addToArrayForm(A, K)
    fmt.Println(ret)
}
