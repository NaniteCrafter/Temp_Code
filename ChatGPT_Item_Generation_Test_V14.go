package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type APIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type RPGItem struct {
	ID                       string                   `json:"id"`
	Name                     string                   `json:"name"`
	Type                     string                   `json:"type"`
	Subtype                  string                   `json:"subtype"`
	Lore                     string                   `json:"lore"`
	SpecialAttributes        []map[string]interface{} `json:"special_attributes"`
	BonusStats               map[string]int           `json:"bonus_stats"`
	Damage                   map[string]interface{}   `json:"damage"`
	ActiveSkills             []map[string]interface{} `json:"active_skills"`
	DispositionTowardsPlayer string                   `json:"disposition_towards_player"`
	Value                    int                      `json:"value"`
	Durability               map[string]int           `json:"durability"`
	Enchantment              map[string]interface{}   `json:"enchantment"`
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set the OPENAI_API_KEY environment variable")
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": `Generate a JSON object for a fantasy RPG item that includes exactly the following fields and no others, in this format:
    {
        "id": "string",
        "name": "string",
        "type": "string",
        "subtype": "string",
        "lore": "string",
        "special_attributes": [{"name": "string", "description": "string"}],
        "bonus_stats": {"strength": "int", "charisma": "int", "agility": "int"},
        "damage": {"amount": "int", "type": "string", "range": "string"},
        "active_skills": [{"name": "string", "description": "string", "mana_cost": "int", "cooldown": "int"}],
        "disposition_towards_player": "string",
        "value": "int",
        "durability": {"current": "int", "max": "int"},
        "enchantment": {"name": "string", "effect": "string", "duration": "string", "cooldown": "string"}
    }`},
		},
	})
	if err != nil {
		log.Fatalf("Error creating request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var apiResponse APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatalf("Error parsing API response: %v", err)
	}

	if len(apiResponse.Choices) > 0 && apiResponse.Choices[0].Message.Content != "" {
		var item RPGItem
		jsonStr := apiResponse.Choices[0].Message.Content
		// Print raw JSON for debugging
		fmt.Println("Raw JSON:", jsonStr)

		if err := json.Unmarshal([]byte(jsonStr), &item); err != nil {
			log.Fatalf("Error parsing item JSON: %v", err)
		}

		fmt.Printf("Item ID: %s\n", item.ID)
		fmt.Printf("Item Name: %s\n", item.Name)
		fmt.Printf("Item Type: %s\n", item.Type)
		// Continue with printing all other attributes...
	} else {
		fmt.Println("No item data received in the response.")
	}
}
