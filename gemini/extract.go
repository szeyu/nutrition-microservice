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
    if err := loadEnv(); err != nil {
        return "", err
    }

    GEMINI_API_KEY := os.Getenv("GEMINI_API_KEY")
    fmt.Println("GEMINI_API_KEY:", GEMINI_API_KEY)

    ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))
    if err != nil {
        return "", err
    }
    defer client.Close()

    tempFile, err := createTempFile(imageData)
    if err != nil {
        return "", err
    }
    defer os.Remove(tempFile.Name())

    file, err := uploadFile(ctx, client, tempFile.Name())
    if err != nil {
        return "", err
    }
    defer client.DeleteFile(ctx, file.Name)

    ingredients, err := generateContent(ctx, client, file.URI)
    if err != nil {
        return "", err
    }

    fmt.Println("Ingredients:", ingredients)
    return ingredients, nil
}

func loadEnv() error {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
        return err
    }
    return nil
}

func createTempFile(imageData string) (*os.File, error) {
    decodedData, err := base64.StdEncoding.DecodeString(imageData)
    if err != nil {
        return nil, err
    }

    fmt.Println("Decoded data size:", len(decodedData))

    tempFile, err := os.CreateTemp("", "*.jpg")
    if err != nil {
        return nil, err
    }

    if _, err := tempFile.Write(decodedData); err != nil {
        return nil, err
    }

    if err := tempFile.Sync(); err != nil {
        return nil, err
    }

    if err := tempFile.Close(); err != nil {
        return nil, err
    }

    fmt.Println("Temp file path:", tempFile.Name())
    return tempFile, nil
}

func uploadFile(ctx context.Context, client *genai.Client, filePath string) (*genai.File, error) {
    tempFile, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer tempFile.Close()

    fileData, err := io.ReadAll(tempFile)
    if err != nil {
        return nil, err
    }

    fmt.Println("File data size:", len(fileData))

    file, err := client.UploadFileFromPath(ctx, filePath, nil)
    if err != nil {
        return nil, err
    }

    return file, nil
}

func generateContent(ctx context.Context, client *genai.Client, fileURI string) (string, error) {
    model := client.GenerativeModel("gemini-1.5-flash")
    resp, err := model.GenerateContent(ctx,
        genai.FileData{URI: fileURI},
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

    return ingredients.String(), nil
}