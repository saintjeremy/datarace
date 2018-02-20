// experiment detecting race conditions in go

package main

import "fmt"

func main() {
	fmt.Println(getNumber())
}

// getNumber function is setting the value of i in a separate goroutine
// also returning i from the function without any knowledge of whether
// our goroutine has completed or not.
// So now, there are two operations that are taking place
// value of i is set to 5, and is also beinf returned from the function

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()

	return i
}
