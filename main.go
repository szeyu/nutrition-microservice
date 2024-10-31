package main

import (
    "encoding/json"
    "net/http"
)

type IngredientRequest struct {
    ImagePath string `json:"image_path"`
}

type NutritionRequest struct {
    Ingredients []string `json:"ingredients"`
}

type NutritionResponse struct {
    NutritionalValues map[string]float64 `json:"nutritional_values"`
}

func extractIngredients(w http.ResponseWriter, r *http.Request) {
    var req IngredientRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call Google Gemini API here and extract ingredients...
    ingredients := []string{"1 cup rice", "10 oz chickpeas"} // Mock response

    json.NewEncoder(w).Encode(ingredients)
}

func analyzeNutrition(w http.ResponseWriter, r *http.Request) {
    var req NutritionRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call Edamam API here and analyze nutrition...
    nutritionResponse := NutritionResponse{
        NutritionalValues: map[string]float64{"calories": 200, "protein": 10}, // Mock response
    }

    json.NewEncoder(w).Encode(nutritionResponse)
}

func main() {
    http.HandleFunc("/extract_ingredients", extractIngredients)
    http.HandleFunc("/analyze_nutrition", analyzeNutrition)

    http.ListenAndServe(":8080", nil)
}
