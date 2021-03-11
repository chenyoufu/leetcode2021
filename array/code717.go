package array

func isOneBitCharacter(bits []int) bool {

    flag := 0
    for i:=0;i < len(bits) -1; i++ {

        if flag == 0 {
            if bits[i] == 1 {
                flag = 1
            }
        } else {
            flag = 0
        }
    }

    if flag == 0 {
        return true
    }
    return false
}
