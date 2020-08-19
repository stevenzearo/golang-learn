package main

func getData() (int32, int32) {
	return 0, 1
}
func main() {
	//	----------variable-----------
	var (
		b bool
		s string
	)
	b = false
	s = "hello, world"
	println(b)
	println(s)
	println("-------------")
	//	----------Anonymous variable-----------
	v1, _ := getData()
	_, v2 := getData()
	println(v1, v2)
	//	----------Data Type------------
	var c1 = 2 + 3i
	var c2 = 1 + 4i
	println(real(c1*c2), imag(c1*c2))
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			println(i)
		}
	}
	var j = 10
	for j >= 0 {
		if j%3 == 0 {
			println(j)
		}
		j--
	}
}
