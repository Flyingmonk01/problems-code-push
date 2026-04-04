func decodeCiphertext(encodedText string, rows int) string {
    cols := len(encodedText) / rows

    mat := make([][]byte, rows)
    it := 0

    for i := 0; i < rows; i++ {
        mat[i] = make([]byte, cols)
        for j := 0; j < cols; j++ {
            mat[i][j] = encodedText[it]
            it++
        }
    }

    result := make([]byte, 0)

    for startCol := 0; startCol < cols; startCol++ {
        i, j := 0, startCol
        for i < rows && j < cols {
            result = append(result, mat[i][j])
            i++
            j++
        }
    }

    end := len(result)
    for end > 0 && result[end-1] == ' ' {
        end--
    }

    return string(result[0:end])
}