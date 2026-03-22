func isEqual(mat [][]int, t [][]int) bool {
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[0]); j++ {
            if mat[i][j] != t[i][j]{
                return false
            }
        }
    }
    return true
}

func rotate(mat [][]int) {
    n := len(mat)
    // => FInding rotation of matrix involves two steps:
    // 1. finding the transpose of matrix
    // 2. reverse the rows for clockwise(90) or reverse the cols for anticlockwise(90)

    for i:=0; i < n; i++ {
        for j := i; j < n; j++ {
            mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
        }
    }

    for i:=0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            mat[i][n-j-1], mat[i][j] = mat[i][j], mat[i][n-j-1]
        }
    }

}


func findRotation(mat [][]int, target [][]int) bool {

    for k := 0; k < 4; k++ {
        if isEqual(mat, target){
            return true
        }
        rotate(mat)
    }

    return false
}