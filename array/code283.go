package array

func moveZeroes(nums []int) {

    var pos = -1
    for i := 0; i < len(nums); i++ {

        if nums[i] == 0 {
            if pos == -1 {
                pos = i
            }
        } else {

            if pos != -1 {
                nums[pos] = nums[i]
                nums[i] = 0
                pos++
            }
        }

    }
}


func moveZeroes2(nums []int) {

    var lastNonZero = 0
    for i := 0; i < len(nums); i++ {

        if nums[i] != 0 {
            nums[lastNonZero] = nums[i]
            lastNonZero++
        }
    }

    for ; lastNonZero < len(nums); lastNonZero++ {
        nums[lastNonZero] = 0
    }
}

func moveZeroes3(nums []int) {

    var lastNonZero = 0
    for i := 0; i < len(nums); i++ {

        if nums[i] != 0 {
            nums[lastNonZero] = nums[i]
            nums[i] = nums[lastNonZero]
            lastNonZero++
        }
    }


}