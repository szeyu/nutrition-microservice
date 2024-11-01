# Go Microservice for Ingredient Extraction and Nutrition Analysis

A microservice in Go to extract ingredient information from an image using Google Gemini API and analyze their nutritional content using Edamam API.

## Prerequisites

- Go 1.16+
- Go modules enabled

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/szeyu/nutrition-microservice.git
   cd nutrition-microservice
   ```

2. Initialize Go modules:
   ```
   go mod init nutrition-microservice
   ```

3. Run the application:
   ```
   go run main.go
   ```
## API Endpoints

- `POST /extract_ingredients`
  - Request body: `{ "image_data": "<base64 encoded image>" }`
  - Response: List of extracted ingredients.

- `POST /analyze_nutrition`
  - Request body: `{ "ingredients": "1 cup rice, 10 oz chickpeas" }`
  - Response: Nutritional values for the ingredients.

## Testing

You can test the endpoints using the provided example scripts:

1. To test ingredient extraction:
   ```
   go run examples/extract_ingredients/test_extract_ingredients.go
   ```

2. To test nutrition analysis:
   ```
   go run examples/analyze_nutrition/test_analyze_nutrition.go
   ```
