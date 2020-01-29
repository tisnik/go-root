package main

func AddFloat32(x float32, y float32, z float32) float32 {
	return x + y + z
}

func AddFloat64(x float64, y float64, z float64) float64 {
	return x + y + z
}

func main() {
	println(AddFloat32(1.0, 2.0, 3.0))
	println(AddFloat64(1.0, 2.0, 3.0))
}
