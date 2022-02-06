package main

func main() {
	loopTillCondition(5)

	loopTillConditionWithPostClause(3)

	infiniteLoop()
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
