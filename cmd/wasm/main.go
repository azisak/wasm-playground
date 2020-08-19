package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

// func prettyJson(input string) (string, error) {
// 	var raw interface{}
// 	if err := json.Unmarshal([]byte(input), &raw); err != nil {
// 		return "", err
// 	}
// 	pretty, err := json.MarshalIndent(raw, "", "  ")
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(pretty), nil
// }

// func jsonWrapper() js.Func {
// 	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		if len(args) != 1 {
// 			return "Invalid no of arguments passed"
// 		}
// 		inputJSON := args[0].String()
// 		fmt.Printf("input %s\n", inputJSON)
// 		pretty, err := prettyJson(inputJSON)
// 		if err != nil {
// 			fmt.Printf("unable to convert to json %s\n", err)
// 			return err.Error()
// 		}
// 		return pretty
// 	})
// 	return jsonFunc
// }

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func firstNPrime(n int) (string, error) {
	if n < 1 {
		return "", errors.New("N should be >= 1")
	}

	result := []string{}
	for i := 2; n > 0; i++ {
		if isPrime(i) {
			result = append(result, strconv.Itoa(i))
			n -= 1
		}
	}

	return strings.Join(result, ","), nil
}

func primeWrapper() js.Func {
	primeFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		N := args[0].Int()
		fmt.Printf("input %s\n", N)
		primes, err := firstNPrime(N)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return err.Error()
		}
		return primes
	})
	return primeFunc
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("firstNPrimeGolang", primeWrapper())
	<-make(chan bool)
}
