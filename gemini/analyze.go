package gemini

import (
    "log"
    "strings"
	"fmt"
	"os"
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
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

	var GEMINI_API_KEY string = os.Getenv("GEMINI_API_KEY")
	fmt.Println("GEMINI_API_KEY:", GEMINI_API_KEY)

	ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

	// Set the model and prompt with image
    model := client.GenerativeModel("gemini-1.5-flash")
    resp, err := model.GenerateContent(ctx,
        genai.Text(
            `Based on the ` + ingredients + `, Give me the nutrition information in JSON format ONLY.
			Don't provide explanation. Please follow the output format.
			Just output with your assumption.
			Don't output backquote character.
            Example output:
            {
				"total_calories": 100,
				"total_protein": 20,
				"total_fat": 10,
				"total_carbs": 20,
				"total_cholesterol": 50,
				"total_sodium": 1000,
				"total_calcium": 10,
				"total_iron": 5,
				"total_magnesium": 100,
				"total_potassium": 200,
			}

            Output format:
            {
				"total_calories": <cal unit>,
				"total_protein": <gram unit>,
				"total_fat": <gram unit>,
				"total_carbs": <gram unit>,
				"total_cholesterol": <gram unit>,
				"total_sodium": <milligram unit>,
				"total_calcium": <milligram unit>,
				"total_iron": <milligram unit>,
				"total_magnesium": <milligram unit>,
				"total_potassium": <milligram unit>,
			}`))

    if err != nil {
        return "", err
    }

	var nutritionJSON strings.Builder
    for _, c := range resp.Candidates {
        if c.Content != nil {
            nutritionJSON.WriteString(fmt.Sprintf("%v", *c.Content))
        }
    }

    fmt.Println("nutritionJSON:", nutritionJSON.String())
    return nutritionJSON.String(), nil
}
