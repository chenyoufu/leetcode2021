package array

func sortArrayByParityII(nums []int) []int {
    var i = 0
    var j = 1

    for {

        if i > len(nums) - 1 || j > len(nums) -1{
            break
        }

        if nums[i]%2 > nums[j]%2 {
            nums[i], nums[j] = nums[j], nums[i]
        }

        if nums[i]%2 == 0 {
            i += 2
        }

        if nums[j]%2 != 0 {
            j += 2
        }

    }

    return nums

}
