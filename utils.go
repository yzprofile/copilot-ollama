package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func dump_req(name string, req *http.Request) {
	if !VERBOSE {
		return
	}
	fmt.Println("Request:", name)
	buf, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
	} else {
		fmt.Println(string(buf))
	}
	fmt.Println("--------------")
}

func dump_rsp(name string, rsp *http.Response) {
	if !VERBOSE {
		return
	}
	fmt.Println("Response:", name)
	buf, err := httputil.DumpResponse(rsp, true)
	if err != nil {
		fmt.Println("Error dumping response:", err)
	} else {
		fmt.Println(string(buf))
	}
	fmt.Println("--------------")
}
