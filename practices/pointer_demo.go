package practices

import "fmt"

func zeroval(ival int) int {
	ival = 0
	return ival
}

func zeroptr(iptr *int) *int {
	*iptr = 0
	return iptr
}

func PointerMain() {
	i := 1
	fmt.Println("initial: ", i)

	fmt.Println("pointer:", &i)

	fmt.Println(zeroval(i))
	fmt.Println(1111)

	fmt.Println("zeroval:", i)

	fmt.Println(zeroptr(&i))
	fmt.Println(2222)

	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}
