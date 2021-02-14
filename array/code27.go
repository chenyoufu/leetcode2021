package array

func removeElement(nums []int, val int) int {
    var pos = 0

    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[pos] = nums[i]
            pos++
        }
    }

    return pos
}
