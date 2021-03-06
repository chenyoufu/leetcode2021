package array

/*
You have a long flowerbed in which some of the plots are planted,
and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's,
where 0 means empty and 1 means not empty, and an integer n,
return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.

Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1
Output: true
Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2
Output: false


Constraints:

1 <= flowerbed.length <= 2 * 104
flowerbed[i] is 0 or 1.
There are no two adjacent flowers in flowerbed.
0 <= n <= flowerbed.length
*/
func canPlaceFlowers(flowerbed []int, n int) bool {

    if n == 0 {
        return true
    }

    if len(flowerbed) == 1 {
        if flowerbed[0] == 0 && n == 1 {
            return true
        }
        return false
    }

    if len(flowerbed) == 2 {
        if flowerbed[0] == 0 && flowerbed[1] == 0 && n == 1 {
            return true
        }
        return false
    }

    if flowerbed[0] == 0 && flowerbed[1] == 0 {
        flowerbed[0] = 1
        n--
    }

    var i = 0
    for ; i < len(flowerbed)-3; i++ {
        if flowerbed[i] != 1 {
            continue
        }

        if flowerbed[i+2] == len(flowerbed) {
        }

        if flowerbed[i+2] == 0 && flowerbed[i+3] == 0 {
            n -= 1
            flowerbed[i+2] = 1
        }
    }

    if len(flowerbed[i:]) == 3 && flowerbed[i] == 1 && flowerbed[i+1] == 0 && flowerbed[i+2] == 0 {
        n -= 1
    }

    if n <= 0 {
        return true
    }

    return false
}

func canPlaceFlowers2(flowerbed []int, n int) bool {

    if n == 0 {
        return true
    }

    for i := 0; i < len(flowerbed); i++ {

        if flowerbed[i] == 0 {
            if i == 0 {
                if len(flowerbed) == 1 || flowerbed[i+1] == 0 {
                    flowerbed[i] = 1
                    n--
                }
            } else if i == len(flowerbed)-1 {
                if flowerbed[i-1] == 0 {
                    flowerbed[i] = 1
                    n--
                }
            } else {
                if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
                    n--
                    flowerbed[i] = 1
                }
            }
        }
    }

    if n <= 0 {
        return true
    }

    return false
}
