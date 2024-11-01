// gemini/extract.go
package gemini

import (
	"context"
    "encoding/base64"
    "log"
    "os"
    "strings"
	"fmt"
	"io"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"github.com/joho/godotenv" // Import godotenv
)

type IngredientRequest struct {
    ImageData string `json:"image_data"`
}

// ExtractIngredients function that uses the Google Gemini API to extract ingredients from an image
func ExtractIngredients(imageData string) (string, error) {
	// Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

	var GEMINI_API_KEY string = os.Getenv("GEMINI_API_KEY")
	fmt.Println("GEMINI_API_KEY:", GEMINI_API_KEY)

	ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Decode base64 image data
    decodedData, err := base64.StdEncoding.DecodeString(imageData)
    if err != nil {
        return "", err
    }

	// Log the size of the decoded data
	fmt.Println("Decoded data size:", len(decodedData))

    // Write the decoded image data to a temporary file
    tempFile, err := os.CreateTemp("", "*.jpg")
    if err != nil {
        return "", err
    }
	
    defer os.Remove(tempFile.Name()) // Clean up the file afterward
	
    // Write data to the temp file and close it
	if _, err := tempFile.Write(decodedData); err != nil {
		return "", err
	}

	// Flush any remaining data to the file
	if err := tempFile.Sync(); err != nil {
		return "", err
	}

	if err := tempFile.Close(); err != nil {
		return "", err
	}

	fmt.Println("Temp file path:", tempFile.Name())

	// Open the temporary file to verify
	tempFile, err = os.Open(tempFile.Name())
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	// Read the contents of the file
	fileData, err := io.ReadAll(tempFile)
	if err != nil {
		return "", err
	}

	// Log the file data size to confirm it was written
	fmt.Println("File data size:", len(fileData))

    // Upload the image file to Gemini API
    file, err := client.UploadFileFromPath(ctx, tempFile.Name(), nil)
    if err != nil {
        return "", err
    }
    defer client.DeleteFile(ctx, file.Name)

    // Set the model and prompt with image
    model := client.GenerativeModel("gemini-1.5-flash")
    resp, err := model.GenerateContent(ctx,
        genai.FileData{URI: file.URI},
        genai.Text(
            `Based on the food scanned, extract out the quanitty and ingredients used in food image. Don't provide explanation. Please follow the output format
            Example output:
            - 4 pieces of Wedges

            Output format:
            - <quantity> <unit> <Ingredient>`))

    if err != nil {
        return "", err
    }

    var ingredients strings.Builder
    for _, c := range resp.Candidates {
        if c.Content != nil {
            ingredients.WriteString(fmt.Sprintf("%v", *c.Content))
        }
    }

    fmt.Println("Ingredients:", ingredients.String())
    return ingredients.String(), nil
}