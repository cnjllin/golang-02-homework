package main

import "strconv"
import "fmt"

func calc(a, b, c string) string {

	a_int, _ := strconv.Atoi(a)
	c_int, _ := strconv.Atoi(c)
	switch b {
	case "*":
		return fmt.Sprintf("%d", a_int*c_int)
	case "/":
		return fmt.Sprintf("%d", a_int/c_int)
	case "+":
		return fmt.Sprintf("%d", a_int+c_int)
	case "-":
		return fmt.Sprintf("%d", a_int-c_int)
	default:
		return "0"
	}

}

func Calc3(a_l []string, b, c string) string {
	//fmt.Printf("%v %v %v\n", a_l, b, c)
	l := len(a_l)
	if l == 1 {
		return calc(a_l[0], b, c)
	}
	return calc(Calc3(a_l[:l-2], a_l[l-2], a_l[l-1]), b, c)
}

func main() {
	a_list := []string{"300", "+", "2", "/", "3"}
	fmt.Println(Calc3(a_list, "/", "3")) // 2
}
