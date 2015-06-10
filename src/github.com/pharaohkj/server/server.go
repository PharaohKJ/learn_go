package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HelloMessage struct {
	Message string
	Time    string
}

func handler(w http.ResponseWriter, r *http.Request) {
	s := HelloMessage{
		Message: "Hello, World",
		Time:    time.Now().Format(time.RFC3339),
	}
	b, err := json.Marshal(s)
	if err == nil {
		fmt.Fprintf(w, string(b))
	}
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}
