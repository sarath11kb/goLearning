package main

import (
	"fmt"
	"learn1/helpers"
	"strings"

	"github.com/pborman/uuid"
)

func main() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)

	name := "harry"
	fmt.Println(name)
	n := 5
	helpers.PrintFirst(n)
	//	helpers.InfiniteCode()
	a, b := 1, 13
	if sum := a + b; sum < 10 {
		fmt.Println(sum)
	} else if sum2 := a + b + 200; sum < 10 {
		fmt.Println("diff is ", sum2)
	} else {
		fmt.Println("sum is ", sum, "sum2 is ", sum2)
	}

	helpers.Defertest()
	helpers.TypeConversion()
	helpers.EvalOrder()
	helpers.ArrayTest()
	var v = helpers.Points{X: 1}
	helpers.SliceTest()
	fmt.Println(v)
	helpers.RangeFunction()

}
