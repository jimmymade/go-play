package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	c1 := make(chan string, 10)
	go addToChan(c1)
	for x := 0; x < 5; x++ {
		value := <-c1
		fmt.Println(value)
	}
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello")

	}
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8081", nil)
}

func addToChan(c1 chan string) {
	for x := 0; x < 5; x++ {
		c1 <- "value: " + strconv.Itoa(x)
	}
	close(c1)
}
