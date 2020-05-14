// HTTP: How to print a raw HTTP response
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func Example() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Date", "Wed, 01 Jan 2020 00:00:00 GMT")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// DumpResponse returns wire representation
	// of the http response
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	// %q is used here to make the testable example work
	// use %s to print in multiple lines:
	// HTTP/1.1 200 OK
	// Content-Length: 11
	// Content-Type: text/plain; charset=utf-8
	// Date: Wed, 01 Jan 2020 00:00:00 GMT
	//
	// hello world
	fmt.Printf("%q", dump)
	// Output:
	// "HTTP/1.1 200 OK\r\nContent-Length: 11\r\nContent-Type: text/plain; charset=utf-8\r\nDate: Wed, 01 Jan 2020 00:00:00 GMT\r\n\r\nhello world"
}
