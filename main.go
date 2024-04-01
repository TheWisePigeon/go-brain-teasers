package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
	"unicode/utf8"
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

// Guessed 6 but it actually prints 7
// Unicode moment, Go strings are UTF-8 encoded, len() returns the size in bytes
// ó is taking 2 bytes hence the length is 7
// We can use RuneCountInString instead to get the number of runes in the string and it will return 6
// A rune is a type alias for int32 and is used to represent a single Unicode character
func krakow() {
	city := "Kraków"
	fmt.Println(len(city))
	fmt.Println(utf8.RuneCountInString(city))
}

// Will not compile
// nil is not a type but a reserved word. If we gave a type to n like `n *int := nil` it would compile
func nil_moment() {
	// Won't compile
	// n := nil
	var n *int = nil
	fmt.Println(n)
}

// Will compile and pring a\tb
// We have 2 types of string literals in Go: interpreted ones enclosed in quotes and raw ones enclosed in back ticks
// If s was an interpreted string the program would print 'a  b' becsue '\t' would be **interpreted** as a tab character
func raw_diet() {
	s := `a\tb`
	fmt.Println(s)
}

// time.Sleep(timeout * time.Millisecond)
// Won't compile because type mismatch, time.Sleep takes in a time.Duration
// Even tho time.Duration is just a type alias for int64, we can not multiply an int and a time.Duration
// We have to cast timeout into a time.Duration
// We can also declare timeout as a const and then it's type will be resolved in the contex of usage
func are_we_there_yet() {
	timeout := 3
	fmt.Print("Before ")
	// time.Sleep(timeout * time.Millisecond)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	fmt.Println("After ")
}

// Yes they can, and they do
// This will print 1.2100000000000002 instead of the correct answer which is 1.21
// Because of the http://en.wikipedia.org/wiki/IEEE_754 flating point numbers specification
// https://docs.oracle.com/cd/E19957-01/806-3568/ncg_goldberg.html
func can_numbers_lie() {
	n := 1.1
	fmt.Println(n * n)
}

// Will output 1,2,3 after the Go team fixed for loops but 2,2,2 in the latter versions of Go
// I don't understand this well enough yet. Need to practice and read more
func sleep_sort() {
	var wg sync.WaitGroup
	for _, n := range []int{3, 1, 2} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond)
			fmt.Printf("%d\n", n)
		}()
	}
	wg.Wait()
	fmt.Println()
}

// Will output false
// To make it simple, when comparing t1 and t2 Go will compare the time.Time struct fields
// By default this struct has a wall clock and a monotonic clock. But when encoding time.Time
// Go doesnt include the monotonic clock reading in the output causing the comparision to fail
// To solve this the Go team recommends using t.Equal when comparing time.Time structs
func just_in_time() {
	t1 := time.Now()
	data, err := json.Marshal(t1)
	if err != nil {
		log.Fatal(err)
	}
	var t2 time.Time
	if err := json.Unmarshal(data, &t2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t1 == t2)
}

// Will out put a=[1 10 3] b=[1 10]
// Because of the way append works, when we do append(a[:1]) it will create a slice of length 1 and of capacity 3
// Because append uses the underlying array, it then checks if there is enough space in the array to add a new element
// In our case we have a capacity of 3 so there is enough space, instead of creating a new array, append will change
// the underlying array and put 10 at the position 1. a and b point to the same underlying array. But why didnt b
// print [1 10 3] too? because b has a length of 2 while a has a length and a capacity of 3
func simple_append() {
	a := []int{1, 2, 3}
	b := append(a[:1], 10)
	fmt.Printf("a=%v b=%v", a, b)
}

func main() {
	simple_append()
}
