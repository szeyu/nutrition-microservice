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
- `POST /edamam_analyze_nutrition`
   - Request body: 
      ```json
      {
         "ingredients": "5 pieces of shawarma, 3 pieces of green pepper, 1 bowl of tomato salad, 1 bowl of sauce, 1 bowl of fries, 1 plate of fries, 1 plate of chicken"
      }
      ```
   - Response: Nutritional values for the ingredients. Kindly refer to [Edamam Output Format](https://api.edamam.com/doc/open-api/nutrition-analysis-v1.yaml)

- `POST /gemini_analyze_nutrition`
   - Request body: 
      ```json
      {
         "ingredients": "5 pieces of shawarma, 3 pieces of green pepper, 1 bowl of tomato salad, 1 bowl of sauce, 1 bowl of fries, 1 plate of fries, 1 plate of chicken"
      }
      ```
   - Response: Nutritional values for the ingredients.
      
      Output format:
      ```json
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
         }
      ```

## Testing

You can test the endpoints using the provided example scripts:

1. To test ingredient extraction:
   ```
   go run examples/extract_ingredients/test_extract_ingredients.go
   ```

2. To test nutrition analysis:

   * For Edamam API:
      ```
      go run examples/edamam_analyze_nutrition/test_edamam_analyze_nutrition.go
      ```

   * For Gemini API:
      ```
      go run examples/gemini_analyze_nutrition/test_gemini_analyze_nutrition.go
      ```
