package main

func Sign(value int) int {
	if value < 0 {
		return -1
	} else if value > 0 {
		return 1
	} else {
		return 0
	}

}

func main() {
	println(Sign(-100))
	println(Sign(100))
	println(Sign(0))
}
