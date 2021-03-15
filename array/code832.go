package array

func flipAndInvertImage(image [][]int) [][]int {

    var (
        n = len(image)
    )
    for i := 0; i < n; i++ {
        row := image[i]
        for j := 0; j < n; j++ {
            if j == n-j-1 {
                row[j] = 1 - row[j]
                break
            }
            if j > n-j-1 {
                break
            }
            row[j], row[n-j-1] = 1-row[n-j-1], 1-row[j]
        }
    }

    return image
}
