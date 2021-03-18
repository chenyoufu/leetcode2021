package array

func validMountainArray(arr []int) bool {

    if len(arr) < 3 {
        return false
    }

    if arr[0] >= arr[1] {
        return false
    }

    isUp := true
    ok := true

    for i := 2; i < len(arr); i++ {
        if arr[i-1] == arr[i] {
            ok = false
            break
        }
        if isUp {
            if arr[i-1] > arr[i] {
                isUp = false
            }
        } else {
            if arr[i-1] < arr[i] {
                ok = false
                break
            }
        }
    }

    return ok && !isUp
}

func validMountainArray2(arr []int) bool {
    i := 0
    // walk up
    for ; i < len(arr)-1 && arr[i] < arr[i+1]; i++ {
    }
    //  peak can't be first or last
    if i == 0 || i == len(arr)-1 {
        return false
    }
    // walk down
    for ; i < len(arr)-1 && arr[i] > arr[i+1]; i++ {
    }
    return i == len(arr)-1
}
