package array

func imageSmoother(M [][]int) [][]int {

    var N  = make([][]int, len(M))

    for i := 0; i < len(M); i++ {
        N[i] = make([]int, len(M[i]))
        for j := 0; j < len(M[i]); j++ {

            count := 1
            sum := M[i][j]

            if i > 0 {
                sum += M[i-1][j]
                count++
            }

            if i < len(M)-1 {
                sum += M[i+1][j]
                count++
            }

            if j > 0 {
                sum += M[i][j-1]
                count++
            }

            if j < len(M[i])-1 {
                sum += M[i][j+1]
                count++
            }

            if i > 0 && j > 0 {
                sum += M[i-1][j-1]
                count++
            }

            if i < len(M)-1 && j < len(M[i])-1 {
                sum += M[i+1][j+1]
                count++
            }

            if i > 0 && j < len(M[i])-1 {
                sum += M[i-1][j+1]
                count++
            }

            if i < len(M)-1 && j > 0 {
                sum += M[i+1][j-1]
                count++
            }

            N[i][j] = sum / count

            // zuo M[i][j-1]
            // you M[i][j+1]
            //shang M[i-1][j]
            //xia M[i+1][j]
            // zuo shang M[i-1][j-1]
            // you shang M[i-1][j+1]
            // zuo xia M[i+1][j-1]
            // you xia M[i+1][j+1]

        }
    }

    return N
}
