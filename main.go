package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    "nutrition-microservice/gemini"
    "nutrition-microservice/edamam"
)

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"status": "ok", "message": "Server is healthy"}
    json.NewEncoder(w).Encode(response)
}

// Handler for extracting ingredients
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

// Handler for Gemini API nutrition analysis
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

// Handler for Edamam API nutrition analysis
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

// Handler for Edamam recipe suggestions
func edamamSuggestRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var req edamam.RecipeRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    recipe, err := edamam.SuggestRecipe(req.Ingredients)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(recipe)
}

func main() {
    // Define routes
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/extract_ingredients", extractIngredientsHandler)
    http.HandleFunc("/gemini_analyze_nutrition", geminiAnalyzeNutritionHandler)
    http.HandleFunc("/edamam_analyze_nutrition", edamamAnalyzeNutritionHandler)
    http.HandleFunc("/edamam_suggest_recipe", edamamSuggestRecipeHandler)

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default to 8080 if PORT is not set
    }
    log.Fatal(http.ListenAndServe(":"+port, nil))

}
