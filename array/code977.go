package array

func sortedSquares(nums []int) []int {
    var res []int

    var mid = len(nums)-1
    for k, v := range nums {
        if v >= 0 {
            mid = k
            break
        }
    }

    i := mid
    j := mid - 1

    for {

        if i == len(nums) || j < 0 {
            break
        }

        if nums[i]*nums[i] <= nums[j]*nums[j] {
            res = append(res, nums[i]*nums[i])
            i++
        } else {

            res = append(res, nums[j]*nums[j])
            j--
        }
    }

    for ; i < len(nums); i++ {
        res = append(res, nums[i]*nums[i])
    }

    for ; j >= 0; j-- {
        res = append(res, nums[j]*nums[j])
    }

    return res
}
