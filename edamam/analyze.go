package edamam

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"io"

	"github.com/joho/godotenv" // Import godotenv
)

type NutritionRequest struct {
	Ingredients string `json:"ingredients"`
}

func AnalyzeNutrition(ingredients string) (string, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return "", err
	}

	var EDAMAM_APP_ID string = os.Getenv("EDAMAM_APP_ID")
	var EDAMAM_APP_KEY string = os.Getenv("EDAMAM_APP_KEY")
	fmt.Println("EDAMAM_APP_ID:", EDAMAM_APP_ID)
	fmt.Println("EDAMAM_APP_KEY:", EDAMAM_APP_KEY)
	fmt.Println("ingredients:", ingredients)

	// Prepare request URL with query parameters
	baseURL := "https://api.edamam.com/api/nutrition-data"
	params := url.Values{}
	params.Add("app_id", EDAMAM_APP_ID)
	params.Add("app_key", EDAMAM_APP_KEY)
	// params.Add("nutrition-type", "logging")
	params.Add("ingr", ingredients)

	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println("requestURL:", requestURL)

	// Send GET request to server
	resp, err := http.Get(requestURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Log raw JSON response for inspection
	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("Raw JSON response:", string(rawData))

	return string(rawData), nil
}
