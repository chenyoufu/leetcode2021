package array

import (
    "reflect"
    "testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            "1",
            args{[]int{4, 3, 2, 7, 8, 2, 3, 1}},
            []int{5, 6},
        },
        {
            "2",
            args{[]int{13, 98, 15, 75, 16, 49, 25, 10, 62, 1, 8, 49, 7, 25, 79, 67, 81, 93, 62, 83, 61, 46, 17, 50, 29, 27, 3, 24, 41, 35, 69, 34, 28, 12, 18, 46, 72, 64, 2, 64, 11, 17, 48, 52, 95, 89, 19, 40, 83, 24, 82, 32, 84, 88, 66, 39, 54, 61, 36, 9, 30, 88, 16, 52, 81, 80, 22, 37, 44, 85, 35, 60, 22, 70, 34, 54, 63, 8, 60, 92, 2, 20, 40, 50, 96, 70, 10, 98, 66, 51, 6, 65, 97, 67, 57, 9, 27, 82, 63, 59}},
            []int{4, 5, 14, 21, 23, 26, 31, 33, 38, 42, 43, 45, 47, 53, 55, 56, 58, 68, 71, 73, 74, 76, 77, 78, 86, 87, 90, 91, 94, 99, 100},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
            }
        })
    }
}
