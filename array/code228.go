package array

import "fmt"

func summaryRanges(nums []int) []string {

    if len(nums) ==  0{
        return []string{}
    }

    if len(nums) == 1 {
        return []string{fmt.Sprintf("%d", nums[0])}
    }

    var ret []string

    var start = 0
    var s string
    for i:=1;i<len(nums);i++ {
        if nums[i] - nums[i-1] > 1 {
            s = fmt.Sprintf("%d->%d", nums[start], nums[i-1])
            if start == i-1 {
                s = fmt.Sprintf("%d", nums[i-1])
            }
            ret = append(ret, s)
            start = i
        }
    }

    if start == len(nums)-1 {
        s = fmt.Sprintf("%d", nums[start])
    } else {
        s = fmt.Sprintf("%d->%d", nums[start], nums[len(nums)-1])
    }

    ret = append(ret, s)

    return ret
}
