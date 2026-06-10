type Car struct {
	position int
	speed    int
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}

	// Pair the position and speed (Fix: Changed 'Cars' to 'Car')
	cars := make([]Car, n)
	for i := 0; i < n; i++ {
		cars[i] = Car{position: position[i], speed: speed[i]}
	}

	// Sort cars by position in descending order
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	// Track fleet arrival times (Fix: Changed 'float' to 'float64')
	var stack []float64

	// Iterate and apply bottleneck logic
	for _, car := range cars {
		time := float64(target-car.position) / float64(car.speed)

		if len(stack) == 0 || time > stack[len(stack)-1] {
			stack = append(stack, time)
		}
	}

	return len(stack)
}