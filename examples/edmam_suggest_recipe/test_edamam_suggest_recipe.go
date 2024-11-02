// test_analyze_nutrition.go
// sends a POST request to the server to analyze the nutrition of a list of ingredients.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	var ingredients string = 
    `
    - chicken breast
    - vegetable oil
    - onion
    - garlic
    - tomato
    - salt
    `

	// Prepare request body
	requestBody, err := json.Marshal(map[string]string{"ingredients": ingredients})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send request to server
	resp, err := http.Post("http://localhost:8080/edamam_suggest_recipe", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var nutritionJSON string
	err = json.NewDecoder(resp.Body).Decode(&nutritionJSON)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("nutritionJSON:", nutritionJSON)
	}
}
