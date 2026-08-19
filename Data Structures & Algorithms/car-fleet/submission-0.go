import (
	"slices"
	"cmp"
)

type Car struct {
    position int
    speed    int
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([]Car, n)
	for i := 0; i < n; i++ {
		cars[i] = Car{
			position: position[i],
			speed: speed[i],
		}
	}

	slices.SortFunc(cars, func(a, b Car) int {
		return cmp.Compare(b.position, a.position)
	})

	stack := []float64{}
	for _, car := range cars {
		time := float64(target - car.position) / float64(car.speed)
		stack = append(stack, time)
		n := len(stack)

		if n >= 2 && stack[n-1] <= stack[n-2] {
			stack = stack[:n-1]
		}
	}
	return len(stack)
}
