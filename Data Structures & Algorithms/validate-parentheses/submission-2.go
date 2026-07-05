func isValid(s string) bool {
    // 括弧の対応をマップで管理
    mapping := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    
    var stack []rune
    
    for _, char := range s {
        // 閉じカッコの場合
        if openBracket, exists := mapping[char]; exists {
            // スタックが空か、対応する開きカッコが一致しない場合
            if len(stack) == 0 || stack[len(stack)-1] != openBracket {
                return false
            }
            // 正しければスタックの最新要素を削除
            stack = stack[:len(stack)-1]
        } else {
            // 開きカッコの場合はスタックに積む
            stack = append(stack, char)
        }
    }
    
    // スタックが空（すべてペアになった）ならtrue
    return len(stack) == 0
}


