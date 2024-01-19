// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	// URL of the API endpoint you want to hit
// 	apiURL := "https://www.nytimes.com/svc/wordle/v2/2024-01-01.json"

// 	// Send a GET request
// 	response, err := http.Get(apiURL)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// Read the response body
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return
// 	}

// 	// Print the response body as a string
// 	fmt.Println("Response:", string(body))

// 	apiURL1 := "https://www.youtube.com/mydummyurl"

// 	// Send a GET request
// 	response1, err1 := http.Get(apiURL1)
// 	if err1 == nil {
// 		fmt.Println(response1)
// 	}

// 	// resp1, err1 := url.ParseRequestURI(apiURL1)
// 	// if err1 != nil {
// 	// 	panic(err1)

// 	// } else {
// 	// 	fmt.Println(resp1)
// 	// }
// }

package main

import (
	"fmt"
	"regexp"
)

func rule1(s string) bool {
	// Atleast five characters must be present
	if len(s) > 5 {
		return true
	} else {
		return false
	}
}

func rule2(s string) bool {
	// Must contain a number in the string
	matched, err := regexp.Match(`[0-9]`, []byte(s))
	if err != nil {
		panic(err)
	}
	return matched
}
func rule3(s string) bool {

	matched, err := regexp.Match(`[A-Z]`, []byte(s))
	if err != nil {
		panic(err)
	}
	return matched
}

func rule4(s string) bool {
	// Special Character
	matched, err := regexp.Match(`[^A-Za-z0-9]`, []byte(s))
	if err != nil {
		panic(err)
	}
	return matched
}
func main() {
	fmt.Println("The Go code is running")
	s := "abc12A%"
	// s1 := "abcdefg"
	// s2 := "abcd1234"
	// s3 := "ABCDEFGH"
	// s4 := "ABCDE123"
	fmt.Println((rule1(s)))
	fmt.Println((rule2(s)))
	fmt.Println((rule3(s)))
	fmt.Println(rule4((s)))

}
