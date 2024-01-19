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
	"sort"
	"strconv"
	"strings"
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

func countRomanNumerals(s string) map[string]int {
	rome := []string{"I", "V", "X", "L", "C", "D", "M"}
	counts := make(map[string]int)

	for _, numeral := range rome {
		counts[numeral] = strings.Count(s, numeral)
	}

	return counts
}

func rule9(s string) bool {
	// Roman numerals must multiply to 35
	values := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	counts := countRomanNumerals(s)

	product := 1
	for numeral, count := range counts {
		product *= values[numeral] ^ count
	}

	return product == 35
}

var romanNumerals = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToDecimal(s string) (int, error) {
	total := 0
	for len(s) > 0 {
		// Sort numerals in descending order to match longer numerals first
		var numerals []string
		for numeral := range romanNumerals {
			numerals = append(numerals, numeral)
		}
		sort.Slice(numerals, func(i, j int) bool {
			return len(numerals[i]) > len(numerals[j])
		})

		// Check for longer numerals first
		for _, numeral := range numerals {
			if len(s) >= len(numeral) && s[:len(numeral)] == numeral {
				total += romanNumerals[numeral]
				s = s[len(numeral):]
				break
			}
		}
	}
	if len(s) > 0 {
		return 0, fmt.Errorf("Invalid Roman numeral: %s", s)
	}
	return total, nil
}

func findRomanNumerals(s string) map[string]int {
	result := make(map[string]int)

	// Define a regular expression to match Roman numerals
	regex := regexp.MustCompile(`(?:` + getRomanNumeralPattern() + `)+`)

	// Find all matches in the input string
	matches := regex.FindAllString(s, -1)

	// Convert Roman numerals to decimal values and store in the result map
	for _, match := range matches {
		value, err := romanToDecimal(match)
		if err == nil {
			result[match] = value
		}
	}

	return result
}

func getRomanNumeralPattern() string {
	pattern := ""
	for numeral := range romanNumerals {
		// Ensure shorter numerals are matched before longer ones
		pattern = numeral + "|" + pattern
	}
	// Trim the trailing "|"
	return pattern[:len(pattern)-1]
}

func main() {
	fmt.Println("The Go code is running")
	s := "abc12A%778marchpepsiVVII"
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

	result := findRomanNumerals(s)

	// Print the results
	for numeral, value := range result {
		fmt.Printf("Roman numerals found: %s = %d\n", numeral, value)
	}

	// fmt.Println(rule5(s))

}
