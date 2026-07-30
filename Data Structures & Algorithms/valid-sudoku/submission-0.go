func isValidSudoku(board [][]byte) bool {
    var rows [9]map[byte]bool
    var cols [9]map[byte]bool
    var boxes [9]map[byte]bool

    for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            val := board[r][c]

            if val == '.' {
                continue
            }

            // 現在のマス (r, c) がどの「3x3ブロック」に属するかを計算（0〜8）
            boxIdx := (r / 3) * 3 + (c / 3)

            if rows[r][val] || cols[c][val] || boxes[boxIdx][val] {
                return false 
            }

            rows[r][val] = true
            cols[c][val] = true
            boxes[boxIdx][val] = true
        }
    }

    return true
}

