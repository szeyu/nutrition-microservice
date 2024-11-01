package main

import (
    "encoding/json"
    "log"
    "net/http"

    "nutrition-microservice/gemini"
    "nutrition-microservice/edamam"
)

func extractIngredientsHandler(w http.ResponseWriter, r *http.Request) {
    var req gemini.IngredientRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    ingredients, err := gemini.ExtractIngredients(req.ImageData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(ingredients)
}

func analyzeNutritionHandler(w http.ResponseWriter, r *http.Request) {
    var req edamam.NutritionRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    nutrition, err := edamam.AnalyzeNutrition(req.Ingredients)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(nutrition)
}

func main() {
    http.HandleFunc("/extract_ingredients", extractIngredientsHandler)
    http.HandleFunc("/analyze_nutrition", analyzeNutritionHandler)

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
