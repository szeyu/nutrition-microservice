# examples/test_analyze_nutrition.py
import argparse
import requests

def test_analyze_nutrition():
    url = "http://127.0.0.1:8000/analyze-nutrition/"
    ingredients = [
        {"name": "rice", "quantity": "1 cup"},
        {"name": "chickpeas", "quantity": "10 oz"}
    ]
    response = requests.post(url, json={"ingredients": ingredients})
    if response.status_code == 200:
        print("Nutritional Analysis:", response.json())
    else:
        print("Failed to analyze nutrition:", response.text)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Test analyze_nutrition endpoint")
    args = parser.parse_args()
    test_analyze_nutrition()
