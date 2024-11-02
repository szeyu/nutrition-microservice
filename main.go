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

func geminiAnalyzeNutritionHandler(w http.ResponseWriter, r *http.Request) {
    var req gemini.NutritionRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    nutrition, err := gemini.AnalyzeNutrition(req.Ingredients)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(nutrition)
}

func edamamAnalyzeNutritionHandler(w http.ResponseWriter, r *http.Request) {
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

func edamamSuggestRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var req edamam.RecipeRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    recipe, err := edamam.SugggestRecipe(req.Ingredients)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(recipe)
}

func main() {
    http.HandleFunc("/extract_ingredients", extractIngredientsHandler)
    http.HandleFunc("/gemini_analyze_nutrition", geminiAnalyzeNutritionHandler)
    http.HandleFunc("/edamam_analyze_nutrition", edamamAnalyzeNutritionHandler)
    http.HandleFunc("/edamam_suggest_recipe", edamamSuggestRecipeHandler)

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
