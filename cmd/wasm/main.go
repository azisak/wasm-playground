package main

import (
	"fmt"
	"syscall/js"
)

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func firstNPrime(n int) {
	if n < 1 {
		fmt.Println("N should be >= 1")
		return
	}

	for i := 2; n > 0; i++ {
		if isPrime(i) {
			fmt.Printf("%d,", i)
			n--
		}
	}
	fmt.Println()
}

func primeWrapper() js.Func {
	primeFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			fmt.Println("Invalid no of arguments passed")
		}
		N := args[0].Int()
		firstNPrime(N)
		return nil
	})
	return primeFunc
}

func main() {
	js.Global().Set("firstNPrimeGolang", primeWrapper())
	<-make(chan bool)
}
