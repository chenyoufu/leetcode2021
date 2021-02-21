package array

func merge(nums1 []int, m int, nums2 []int, n int) {

    var nums = make([]int, 0)

    i, j := 0, 0
    for i < m && j < n {
        if nums1[i] <= nums2[j] {
            nums = append(nums, nums1[i])
            i++
        } else {
            nums = append(nums, nums2[j])
            j++
        }
    }

    if i < m {
        nums = append(nums, nums1[i:]...)
    }

    if j < n {
        nums = append(nums, nums2[j:]...)
    }

    for i := 0; i < m+n; i++ {
        nums1[i] = nums[i]
    }

    return
}

func merge2(nums1 []int, m int, nums2 []int, n int) {

    i, j, k := m-1, n-1, m+n-1
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
        k--
    }

    /* 这块可以不需要
       for ;i >= 0;i-- {
           nums1[k] = nums1[i]
           k--
       }
    */

    for ; j >= 0; j-- {
        nums1[k] = nums2[j]
        k--
    }

    return
}
