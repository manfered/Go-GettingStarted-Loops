package main

func main() {
	loopTillCondition(5)
}

// while equivalent loop
func loopTillCondition(max int) {
	var index int
	for index < max {
		println(index)
		index++
	}
}
