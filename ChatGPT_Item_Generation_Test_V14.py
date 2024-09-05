import json
import requests

class APIResponse:
    def __init__(self, choices):
        self.choices = [choice["message"]["content"] for choice in choices]

class RPGItem:
    def __init__(self, item_data):
        self.id = item_data["id"]
        self.name = item_data["name"]
        self.type = item_data["type"]
        self.subtype = item_data["subtype"]
        self.lore = item_data["lore"]
        self.special_attributes = item_data["special_attributes"]
        self.bonus_stats = item_data["bonus_stats"]
        self.damage = item_data["damage"]
        self.active_skills = item_data["active_skills"]
        self.disposition_towards_player = item_data["disposition_towards_player"]
        self.value = item_data["value"]

def fetch_api_response(url, headers, payload):
    try:
        response = requests.post(url, headers=headers, json=payload)
        response.raise_for_status()
        return APIResponse(response.json()["choices"])
    except requests.exceptions.RequestException as e:
        print(f"API request failed: {e}")
        return None

# Example usage
url = "https://api.example.com/items"
headers = {"Authorization": "Bearer your_token"}
payload = {"item_id": "123"}

api_response = fetch_api_response(url, headers, payload)

if api_response:
    for choice in api_response.choices:
        print(choice)
