package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func fetcher()string{
	data := []byte("No data")
	rsp, err := http.Get("http://127.0.0.1:5000/ind")
	if err != nil {
		fmt.Println("failed with error", err)
	} else {
		data, _ = ioutil.ReadAll(rsp.Body)
	}
	return string(data)
}

func main() {
	fmt.Println("Entered the new application..")
	//data := []string("hello")
	data := "Nothing" 

//	data = go fetcher()
//	fmt.Println(data)
//	data = go fetcher()
//	fmt.Println(data)
//	data = fetcher()
//	fmt.Println(data)
//	data = fetcher()
//	fmt.Println(data)

	go fetcher()
	fmt.Println(data)
}
