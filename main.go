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
)

func rule1(s string) bool {

}

func main() {
	fmt.Println("The Go code is running")

}
