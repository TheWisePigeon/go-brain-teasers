package main

import (
	"fmt"
)

// Will compile because π is a valid identifier
// also π being a constant, the 22 will be converted to a float by the compiler before the division
// if we defined lets say a,b to contain 22 and 7.0 respectively, computing a/b would not compile
func pi() {
	var π = 22 / 7.0
	fmt.Println(π)
}

// Will compile, will print 0 because classic Go 0 value moment
// We can access the length of a nil map and also access the value of a key in a nil map
// without causing a panic
func empty_handed() {
	var m map[string]int
	fmt.Println(m["errors"])
}

func main() {

}
