package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	servers := []string{
		"https://uma-v2.appspot.com",
		"https://triage-dot-uma-v2.appspot.com",
		"https://nodeserver-dot-uma-v2.appspot.com",
		"https://gps-dot-uma-v2.appspot.com/v1/users",
		"https://bigquery-dot-uma-v2.uc.r.appspot.com/",
		"https://pol-dot-uma-v2.appspot.com",
		"https://providers-dot-uma-v2.appspot.com",
	}
	for _, server := range servers {
		checkServers(server)
	}
	elapsed := time.Since(start)
	fmt.Printf("Executing time %s", elapsed)
}

func checkServers(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "Server down :(")
	} else {
		fmt.Println(server, "Server ok :)")
	}
}
