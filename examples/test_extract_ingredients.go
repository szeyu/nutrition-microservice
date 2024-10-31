package main

import (
    "bytes"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    imagePath := "C:\\Users\\szeyu\\Downloads\\WhatsApp Image 2024-08-30 at 13.54.18_28b332e2.jpg"
    
    // Read image file
    imageData, err := ioutil.ReadFile(imagePath)
    if err != nil {
        fmt.Println("Error reading image:", err)
        return
    }

    // Check the size of the image data
    fmt.Println("Image read successfully, size:", len(imageData))

    // Convert image data to base64
    encodedImage := base64.StdEncoding.EncodeToString(imageData)

    // Check the size of the encoded image data
    fmt.Println("Encoded image size:", len(encodedImage))

    // Prepare request body
    requestBody, err := json.Marshal(map[string]string{"image_data": encodedImage})
    if err != nil {
        fmt.Println("Error marshalling JSON:", err)
        return
    }

    // Send request to server
    resp, err := http.Post("http://localhost:8080/extract_ingredients", "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    var ingredients string
    err = json.NewDecoder(resp.Body).Decode(&ingredients)
    if err != nil {
        fmt.Println("Error decoding response:", err)
        return
    }
    fmt.Println("Extracted Ingredients:", ingredients)
}
