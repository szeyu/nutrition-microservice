package edamam

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"log"
	"net/url"
	"github.com/joho/godotenv" // Import godotenv
)

type RecipeRequest struct {
	Ingredients string `json:"ingredients"`
}

func SugggestRecipe(ingredients string) (string, error) {
	// Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

	var EDAMAM_APP_ID string = os.Getenv("EDAMAM_RECIPE_APP_ID")
	var EDAMAM_APP_KEY string = os.Getenv("EDAMAM_RECIPE_APP_KEY")
	fmt.Println("EDAMAM_RECIPE_APP_ID:", EDAMAM_APP_ID)
	fmt.Println("EDAMAM_RECIPE_APP_KEY:", EDAMAM_APP_KEY)
	fmt.Println("ingredients:", ingredients)

	// Prepare request URL with query parameters
	baseURL := "https://api.edamam.com/api/recipes/v2"
	params := url.Values{}
	params.Add("app_id", EDAMAM_APP_ID)
	params.Add("app_key", EDAMAM_APP_KEY)
	params.Add("type", "public")
	params.Add("q", ingredients)

	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

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

	fmt.Println("requestURL:", requestURL)

	return string(rawData), nil
}