package array

func majorityElement(nums []int) int {

    var count int
    var maj = nums[0]
    for _, v := range nums {

        if count == 0{
            maj = v
        }

        if maj == v {
            count+=1
        }else {
            count-=1
        }


    }

    return maj
}
