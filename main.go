package main

import (
	"fmt"
	"io"
	"net/http"
)

type testPage struct {
	Stuff string `json:"items"`
}

func main() {

	res, err := http.Get("https://myclu.callutheran.edu/health-check")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("%s", body)
}
