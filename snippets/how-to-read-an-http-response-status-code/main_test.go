// HTTP: How to read an HTTP response status code
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func Example() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Printf("%s\n", resp.Status)
	fmt.Printf("%d\n", resp.StatusCode)
	// Output:
	// 200 OK
	// 200
}
