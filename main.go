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

// Will output the date but not the Hello
// Because we embedded time.Time into the Log struct so it has all the methods and fields of time.Time
// time.Time has a String() string method which means it implements fmt.Stringer interface. fmt.Printf will then
// use the String() method instead of the default output
func what_da_log_doin() {
	type Log struct {
		Message string
		time.Time
	}
	ts := time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC)
	log := &Log{"Hello", ts}
	fmt.Printf("%v\n", log)
}

// Outputs 0.25
// Because in Go this is a hexadecimal floating-point literal. To calculate the value we do
// The value of the literal before the p which is 0x1 in hexadecimal which is 1
// 2 to the power of the value after the p which is 2^-2 which is also 1/2^2 : 0.25
// And we multiply the two values 1x0.25 is 0.25
// https://go.dev/ref/spec#Lexical_elements
func funky_number() {
	fmt.Println(0x1p-2)
}

func fibs(n int) chan int {
	ch := make(chan int)
	go func() {
		a, b := 1, 1
		for i := 0; i < n; i++ {
			ch <- a
			a, b = b, a+b
		}
	}()
	return ch
}

// Dont understand yet. Need to read more bout concurrency patterns in Go
func free_range_ints() {
	for i := range fibs(5) {
		fmt.Printf("%d\n", i)
	}
}

// Will output 1,3,4
// b will be reassigned to because it is an existing variable
func who_is_you() {
	a, b := 1, 2
	b, c := 3, 4
	fmt.Println(a, b, c)
}

// Will output false
// The two strings look the same but are not the same at the byte level
// The first contain the special character ó but the second contains o followed by a control character
func two_cities() {
	city1, city2 := "Kraków", "Kraków"
	fmt.Println(city1 == city2)
}

// Outputs 2, 0
// We create a buffered channel of capacity 2
// We then put 1, and 2 in the channel. Read from it and close it
// When we try to read from a closed channel, if the channel contains a value we will receive it, if not
// We get the 0 value of the channel type
// Here we read once from the channel before so 1 is gone, we then read 2 into a and 0 into b because
// the channel is empty after we read 2
func what_is_in_chanel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	<-ch
	close(ch)
	a := <-ch
	b := <-ch
	fmt.Println(a, b)
}

// Outputs ©
// string is the set of all strings of 8-bit bytes, conventionally but not necessarily representing UTF-8-encoded text
// Which means that doing string(169) will output the unicode representation of 169 which is ©
// To convert 169 into a string literal we can either use strconv.Itao or fmt.Sprintf(i)
func int64resting() {
	i := 169
	// s := string(i)
	s := fmt.Sprintf("%d", i)
	fmt.Println(s)
}

type Job struct {
	State string
	done  chan struct{}
}

func (j *Job) Wait() {
	<-j.done
}

func (j *Job) Done() {
	j.State = "done"
	close(j.done)
}

// Dont understand well concurrency patterns in Go
// Skill issue ;)
func job() {
	ch := make(chan Job)
	go func() {
		j := <-ch
		j.Done()
	}()
	job := Job{"ready", make(chan struct{})}
	ch <- job
	job.Wait()
	fmt.Println(job.State)
}

type OSError int

func (e *OSError) Error() string {
	return fmt.Sprintf("error #%d", *e)
}

func FilePathExists(path string) (bool, error) {
	var err *OSError
	return false, err
}

// Outputs error: nil
// Dont understand well
// Need to read more
func err_or_not_err() {
	if _, err := FilePathExists("/no/such/file"); err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Println("OK")
	}
}

func what_in_da_string() {
	msg := "π = 3.14159265358..."
	fmt.Printf("%T ", msg[0])
	for _, c := range msg {
		fmt.Printf("%T\n", c)
		break
	}
}

func init() {
	fmt.Println("A")
}

func init() {
	fmt.Println("B")
}

func count_me_in() {
	var count int
	var wg sync.WaitGroup
	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// func main() {
// 	count_me_in()
// }
