package array

//import "math"

func plusOne(digits []int) []int {
    var add1 = true

    for i := len(digits) - 1; i >= 1; i-- {
        //x = x + digits[i] * int(math.Pow10(len(digits)-i-1))
        if add1 {
            digits[i]++
            if digits[i] == 10 {
                digits[i] = 0
                add1 = true
            } else {
                add1 = false
                break
            }
        }
    }

    var newDigits []int

    if add1 {
        digits[0]++
        if digits[0] == 10 {
            digits[0] = 0
            newDigits = append(newDigits, 1)
            newDigits = append(newDigits, digits...)
        } else {
            newDigits = digits
        }
    }else {
        newDigits = digits
    }

    return newDigits
}
