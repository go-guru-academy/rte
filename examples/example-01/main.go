package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-guru-academy/rte"
)

// Custom input to be passed down the request chain
type MyInput struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("example-01: started")
	r, _ := rte.New(rte.GORILLA)
	r.Get("/example", example, &MyInput{}, []rte.Middleware{
		middleware01,
		middleware02,
	})
	if err := http.ListenAndServe(":8888", r); err != nil {
		fmt.Println(err)
		return
	}
}

func middleware01(next func(*rte.Default, interface{})) rte.Handler {
	fmt.Println("middleware01: appended to middleware chain")
	return func(defaultInput *rte.Default, input interface{}) {
		fmt.Println("middleware01: pre-request")
		i := input.(*MyInput)
		i.Name = "eric"
		next(defaultInput, i)
		fmt.Println("middleware01: post-request")
	}
}

func middleware02(next func(*rte.Default, interface{})) rte.Handler {
	fmt.Println("middleware02: appended to middleware chain")
	return func(defaultInput *rte.Default, input interface{}) {
		fmt.Println("middleware: pre-request")
		i := input.(*MyInput)
		i.Age = 36
		next(defaultInput, i)
		fmt.Println("middleware: post-request")
	}
}

func example(defaultInput *rte.Default, input interface{}) {
	fmt.Println("status: request received")

	// Cast input to the correct type
	i := input.(*MyInput)

	fmt.Printf("MyInput: %+v\n", i)

	// Marshal the input into JSON
	response, _ := json.Marshal(i)

	// Set the response headers
	defaultInput.W.WriteHeader(http.StatusOK)
	defaultInput.W.Header().Set("Content-Type", "application/json")

	// Write the response
	defaultInput.W.Write(response)
}
