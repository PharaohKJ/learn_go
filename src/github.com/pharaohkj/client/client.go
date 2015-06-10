package main

import (
	//	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//	"time"
)

func main() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer resp.Body.Close()
		contents, re := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", re)
			return
		}
		fmt.Printf("%s\n", string(contents))
	}
}
