package main

//import "strings"
import "fmt"

// Returns a new slice containing the results of applying
// the function `f` to each string in the original slice.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// 全部转换成大写，并且统计字母总数
func mapper(s string) string {
	return
}

func main() {

	// Here we try out our various collection functions.
	var strs = []string{"peach", "apple", "pear", "plum"}

	// The above examples all used anonymous functions,
	// but you can also use named functions of the correct
	// type.
	fmt.Println(Map(strs, mapper))

}
