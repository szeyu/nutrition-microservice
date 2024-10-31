# FastAPI Microservice for Ingredient Extraction and Nutrition Analysis
A microservice in Python to extract ingredients information from image or suggest recipes using Gemini API and Edamam API. Also returns the nutritional content of the food.

This microservice provides two main endpoints for extracting ingredients from an image and analyzing their nutritional value.

## Prerequisites

- Python 3.7+
- FastAPI
- Uvicorn
- Requests (for testing scripts)

## Installation

1. Clone the repository:
    ```
    git clone https://github.com/szeyu/nutrition-microservice.git
    cd nutrition-microservice
    ```
2. Install dependencies:
    `pip install -r requirements.txt`

3. Run the FastAPI application:
    `uvicorn app:app --reload`

The app should now be running at http://127.0.0.1:8000.

## Endpoints

### 1. POST /extract-ingredients/

- **Description**: Accepts an image and extracts the ingredients and their quantities.
- **Payload**: An image file.
- **Response**: A JSON object with extracted ingredients and quantities.

### 2. POST /analyze-nutrition/

- **Description**: Accepts a list of ingredients and returns nutritional information.
- **Payload**: A JSON object containing `ingredients`, a list of objects with `name` and `quantity` fields.
- **Response**: A JSON object with nutritional information.

## Testing

You can use the provided CLI scripts to test each endpoint.

### Test `extract_ingredients` Endpoint

    python examples/test_extract_ingredients.py --image_path "path/to/your/image.png"

### Test `analyze_nutrition` Endpoint

    python examples/test_analyze_nutrition.py

## API Documentation

FastAPI automatically generates API documentation, accessible at:
- **Swagger UI**: http://127.0.0.1:8000/docs
- **ReDoc**: http://127.0.0.1:8000/redoc
