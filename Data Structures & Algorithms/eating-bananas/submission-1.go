func minEatingSpeed(piles []int, h int) int {
	r := 0
	l := 1
	// pilesの最大値が最速で食べ切れる
	for _, p := range piles {
		if p > r {
			r = p
		}
	}
	res := r

	// 速度を二分探索
	for l <= r {
		k := l + (r-l)/2
		totalTime := 0

		for _, p := range piles {
			totalTime += int(math.Ceil(float64(p) / float64(k)))
		}

		if totalTime <= h {
			res = k
			r = k-1
		} else {
			l = k+1
		}
	}
	return res
}
