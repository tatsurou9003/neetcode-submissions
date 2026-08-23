func largestRectangleArea(heights []int) int {
	maxArea := 0
	
	hcopy := make([]int, len(heights), len(heights)+1)
	copy(hcopy, heights)
	hcopy = append(hcopy, 0)

	stack := []int{}

	for i := 0; i < len(hcopy); i++ {
		for len(stack) > 0 && hcopy[stack[len(stack)-1]] > hcopy[i] {
			topIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := hcopy[topIdx]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			 
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}
	return maxArea
}
