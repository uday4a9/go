package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

var val int = 0

func fetcher(c chan string) {
	data := []byte("No data")
	rsp, err := http.Get("http://127.0.0.1:5000/ind")
	if err != nil {
		fmt.Println("failed with error", err)
	} else {
		data, _ = ioutil.ReadAll(rsp.Body)
	}
	val += 1

	if val %2 == 0 {
		time.Sleep(time.Second * 1)
	}

	c <- string(data) + strconv.Itoa(val)
}

func main() {
	fmt.Println("Entered the new application..")

	ch := make(chan string)

	go fetcher(ch)
	go fetcher(ch)
	go fetcher(ch)
	go fetcher(ch)
	fmt.Println(<-ch, <-ch, <-ch, <-ch)

	// This below line cause for indefinite block
	// until osmebody write response to the channel
	//fmt.Println(<-ch, <-ch, <-ch, <-ch)

}
