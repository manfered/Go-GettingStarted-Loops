package main

func main() {
	loopTillCondition(5)

	loopTillConditionWithPostClause(3)

	infiniteLoop()

	loopOverCollection()
}

// while equivalent loop
func loopTillCondition(max int) {
	var index int
	for index < max {
		println(index)
		index++
	}
}

// typical for loop
func loopTillConditionWithPostClause(max int) {
	for index := 0; index < max; index++ {
		println(index)
	}
}

func infiniteLoop() {
	var index int
	for {
		if index == 5 {
			break
		}
		println(index)
		index++
	}
}

// for each loop equivalent
func loopOverCollection() {
	testSlice := []int{1, 2, 3, 4}

	// grabbing the index and value
	for index, value := range testSlice {
		println(index, value)
	}

	// grabbing index only
	for index := range testSlice {
		println(index)
	}

	// grabbing value only by using WriteOnly variable for index
	for _, value := range testSlice {
		println(value)
	}

	// loop through map data type
	testMap := map[string]int{"Fred": 1, "Manfered": 2}
	for key, value := range testMap {
		println(key, value)
	}
}
