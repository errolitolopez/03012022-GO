package main

import "fmt"

func main() {

	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Printf("websites: %v\n", websites)
	fmt.Printf("websites[\"Google\"]: %v\n", websites["Google"])

	keys := make([]string, 0, len(websites))
	for k := range websites {
		keys = append(keys, k)
	}
	fmt.Printf("keys: %v\n", keys)

	for k, v := range websites {
		fmt.Printf("k: %v v: %v\n", k, v)
	}

	websites["Yahoo"] = "https://yahoo.com"

	fmt.Printf("websites: %v\n", websites)

	delete(websites, "Google")

	fmt.Printf("websites: %v\n", websites)
}
