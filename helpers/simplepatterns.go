package helpers

import "fmt"

func PrintFirst(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i, "")
	}
}

func InfiniteCode() {
	for {
	}
}

func Defertest() {

	defer func() {
		fmt.Println("This is a defer func that is created on top of this function")
		fmt.Println("This is the second line in the defer function")
	}()
	fmt.Println("testin defer")
	var i = 0
	for ; i < 10; i++ {
		// calculated and stored in the stack in the order the statement is given
		defer fmt.Println("defer ", i)
	}
	i += 1

	if i < 10 {
		// wont execute because defer is calculated and stored in the stack
		fmt.Println("before ending i = ", i)
	}
	fmt.Println("defer completed with i = ", i)
}

func DeferTest2() {

}

func getSum(a, b int) {
	var sum = (1 + a + b + int(32.0)*4/9)
	//	var sum = (1 + a + b) * 40.00 / 9

	fmt.Printf("type is %T", sum)
}

func EvalOrder() {
	fmt.Println()
	fmt.Println(1 / 3 * 4)
	fmt.Println(4 + 1/3)

}

func TypeConversion() {
	x := 72
	fmt.Println("int value ", x)
	//make y float and it will give error that getsum needs float
	y := 43
	getSum(y, y)

}

func Variables() {
	name, age := "yelan", 23
	fmt.Println(name, " ", age)
}
func StringFunctions() {

}
