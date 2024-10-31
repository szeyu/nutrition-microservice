# examples/test_extract_ingredients.py
import argparse
import requests

def test_extract_ingredients(image_path):
    url = "http://127.0.0.1:8000/extract-ingredients/"
    with open(image_path, "rb") as image_file:
        response = requests.post(url, files={"image_data": image_file})
    if response.status_code == 200:
        print("Extracted Ingredients:", response.json())
    else:
        print("Failed to extract ingredients:", response.text)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Test extract_ingredients endpoint")
    parser.add_argument("--image_path", required=True, help="Path to the image file")
    args = parser.parse_args()
    test_extract_ingredients(args.image_path)
