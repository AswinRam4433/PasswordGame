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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func rule5(s string) bool {
	// Digits must add to 25
	regex := regexp.MustCompile(`\d`)
	indices := regex.FindAllStringIndex(s, -1)
	sum := 0
	for _, indexPair := range indices {
		startIndex := indexPair[0]
		endIndex := indexPair[1]
		i, err := strconv.Atoi(s[startIndex:endIndex])
		if err != nil {
			panic(err)
		}
		sum = sum + i
	}
	return sum == 25

}

func rule6(s string) bool {
	// Contains a month
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	for i := 0; i < len(months); i++ {
		matched, err := regexp.Match(months[i], []byte(s))
		if err != nil {
			panic(err)
		}
		if matched == true {
			return true
		}
	}
	for i := 0; i < len(months); i++ {
		matched, err := regexp.Match(strings.ToLower(months[i]), []byte(s))
		if err != nil {
			panic(err)
		}
		if matched == true {
			return true
		}
	}
	return false
}

func rule7(s string) bool {
	rome := []string{"I", "V", "X", "L", "C", "D", "M"}
	for i := 0; i < len(rome); i++ {
		matched, err := regexp.Match(rome[i], []byte(s))
		if err != nil {
			panic(err)
		}
		if matched == true {
			return true
		}
	}
	return false

}

func rule8(s string) bool {
	// One sponsor: Pepsi, Starbucks, Shell
	sponsors := []string{"Pepsi", "Starbucks", "Shell"}
	for i := 0; i < len(sponsors); i++ {
		matched, err := regexp.Match(sponsors[i], []byte(s))
		if err != nil {
			panic(err)
		}
		if matched == true {
			return true
		}
	}
	for i := 0; i < len(sponsors); i++ {
		matched, err := regexp.Match(strings.ToLower(sponsors[i]), []byte(s))
		if err != nil {
			panic(err)
		}
		if matched == true {
			return true
		}
	}
	return false

}

// func rule9(s string) bool {
// 	// Roman Numerals Multiply to 35
// 	pat1 := `^[A-Za-z0-9]*V[A-Za-z0-9]*VII[A-Za-z0-9]$`
// 	matched, err := regexp.Match(pat1, []byte(s))
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println((matched))
// 	return true

// }

func rule9(s string) bool {
	// Roman Numerals Multiply to 35
	pat1 := `^.*V.*VII.*$`
	matched, err := regexp.Match(pat1, []byte(s))
	if err != nil && strings.Count(s, "V") == 2 && strings.Count(s, "I") == 2 {
		panic(err)
	}
	if matched == true {
		return true
	}

	pat2 := `^.*VII.*V.*$`
	matched1, err1 := regexp.Match(pat2, []byte(s))
	if err1 != nil {
		panic(err1)
	}
	if matched1 == true && strings.Count(s, "V") == 2 && strings.Count(s, "I") == 2 {
		return true
	}

	pat3 := `^.*XXXV.*$`
	matched2, err2 := regexp.Match(pat3, []byte(s))
	if err2 != nil {
		panic(err1)
	}
	if matched2 == true && strings.Count(s, "V") == 1 && strings.Count(s, "X") == 3 {
		return true
	}

	pat4 := `^.*XXXV.*I.*$`
	matched3, err3 := regexp.Match(pat4, []byte(s))
	if err3 != nil {
		panic(err1)
	}
	if matched3 == true && strings.Count(s, "X") == 3 && strings.Count(s, "I") == 1 && strings.Count(s, "V") == 1 {
		return true
	}

	pat5 := `^.*I.*XXXV.*$`
	matched4, err4 := regexp.Match(pat5, []byte(s))
	if err4 != nil {
		panic(err1)
	}
	if matched4 == true && strings.Count(s, "X") == 3 && strings.Count(s, "I") == 1 && strings.Count(s, "V") == 1 {
		return true
	}

	return false
}

type Response struct {
	ID              int    `json:"id"`
	Solution        string `json:"solution"`
	PrintDate       string `json:"print_date"`
	DaysSinceLaunch int    `json:"days_since_launch"`
	Editor          string `json:"editor"`
}

func rule11(s string) bool {
	var resp Response
	apiURL := "https://www.nytimes.com/svc/wordle/v2/2024-01-01.json"

	// Send a GET request
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)

	}
	// defer response.Body.Close()
	// fmt.Println(response)

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)

	}

	// Unmarshal the JSON string into the Response struct
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	// Access the "solution" field
	solution := resp.Solution

	// Print the value of the "solution" field
	fmt.Println("Solution:", solution)
	matched, err := regexp.Match(solution, []byte(s))
	if err != nil {
		panic(err)
	}
	return matched

}

func rule12(s string) bool {
	// Moon Phase
	// From https://jivebay.com/calculating-the-moon-phase/
	round := func(x float64) float64 {
		if x < 0.0 {
			return float64(int(x - 0.5))
		}
		return float64(int(x + 0.5))
	}

	timestamp := time.Now()
	year, month, day := timestamp.Year(), int(timestamp.Month()), timestamp.Day()

	// Moon phase calculation
	var c, e, jd, b float64

	if month < 3 {
		year--
		month += 12
	}

	month++
	c = 365.25 * float64(year)
	e = 30.6 * float64(month)
	jd = c + e + float64(day) - 694039.09
	jd /= 29.5305882
	b = jd
	jd -= b
	b = round(jd * 8)

	if b >= 8 {
		b = 0
	}

	moon_phase := "-1"
	switch int(b) {
	case 0:
		moon_phase = "🌑"
	case 1:
		moon_phase = "🌒"
	case 2:
		moon_phase = "🌓"
	case 3:
		moon_phase = "🌔"
	case 4:
		moon_phase = "🌕"
	case 5:
		moon_phase = "🌖"
	case 6:
		moon_phase = "🌗"
	case 7:
		moon_phase = "🌘"
	default:
		panic("Error In Moon Phase Calculation")
	}
	// fmt.Println(moon_phase)
	matched, err := regexp.Match(moon_phase, []byte(s))
	if err != nil {
		panic(err)
	}
	if matched == true {
		return true
	}

	return false

}

func main() {
	fmt.Println("The Go code is running")
	s := "abc12A778marchpepsiXXXVjdlksImural🌑"
	// s1 := "abcdefg"
	// s2 := "abcd1234"
	// s3 := "ABCDEFGH"
	// s4 := "ABCDE123"

	// fmt.Println((rule1(s)))
	// fmt.Println((rule2(s)))
	// fmt.Println((rule3(s)))
	// fmt.Println((rule4(s)))
	// fmt.Println((rule5(s)))
	// fmt.Println(rule6(s))
	// fmt.Println(rule7(s))
	// fmt.Println(rule8(s))
	// fmt.Println(rule9(s))
	// fmt.Println(rule10(s))

	// fmt.Println(rule5(s))
	// fmt.Println(rule11(s))
	fmt.Println(rule12(s))

	// Print the response body as a string
	// fmt.Println("Response:", string(body["solution"]))

}
