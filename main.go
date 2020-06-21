package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// go sexyCount("trea")
	// sexyCount("bear")
	go sexyCount("trea")
	go sexyCount("bear")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

var errRequestFaild = errors.New("Request failed")

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFaild
	}
	return nil
}
