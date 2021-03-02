package array

func findDisappearedNumbers(nums []int) []int {
    var ret []int

    for i:=0;i < len(nums);i++ {

        v := nums[i]
        if v < 0 {
            v = -v
        }

        if nums[v-1] > 0 {
            nums[v-1] = -nums[v-1]
        }

    }

    for i:=0;i < len(nums);i++ {
        if nums[i] > 0 {
            ret = append(ret, i+1)
        }
    }

    return ret
}
