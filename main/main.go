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
	println("-------------")
	//	----------Data Type------------
	var c1 = 2 + 3i
	var c2 = 1 + 4i
	println(real(c1*c2), imag(c1*c2))
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			println(i)
		}
	}
	println("-------------")
	// ------------goto statement---------------
	var j = 10
	for {
		if j%3 == 0 {
			println(j)
		}
		if j == 0 {
			goto gotoHere
		}
		j--
	}
gotoHere:
	println("done")

	println("-------------")
	// -------------break statement---------------------
	var k = 10
	for {
		if k == 0 {
			break
		}
		println("k:", k)
		k--
	}
	var x = 10
breakStatement:
	for {
		if x == 0 {
			break breakStatement
		}
		println("x:", x)
		x--
	}

	println("-------------")
	// -------continue statement------------------
	var y = 10
	for {
		if y%2 == 0 {
			y--
			continue
		} else {
			println("y:", y)
		}
		y--
		if y == 0 {
			break
		}
	}

continueStatement:
	for i := 1; i < 100; i++ {
		for j := 1; j <= i; j++ {
			if i <= 9 {
				print(j, "*", i, " = ", i*j)
				print("\t\t\t")
			} else {
				println("continuing")
				continue continueStatement
			}
			if j == 10 {
				break continueStatement
			}
		}
		println()
	}
}
