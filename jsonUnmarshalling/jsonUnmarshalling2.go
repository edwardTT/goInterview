package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Status int `json:"status"` //单引号的字符串为标签tag
}

func main() {
	var data = []byte(`{"status": 200}`)
	result := &Result{}

	if err := json.Unmarshal(data, result); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("result=%+v", result)
}
