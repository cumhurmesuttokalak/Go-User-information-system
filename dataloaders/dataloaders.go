package dataloaders

import (
	"encoding/json"

	. "../models"
	util "../utils"
)

func LoadUsers() []User {
	bytes, _ := util.ReadFile("../json/users.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func LoadInterests() []Interest {
	bytes, _ := util.ReadFile("../json/interests.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func LoadInterestMappings() []InteresetMapping {
	bytes, _ := util.ReadFile("../json/userInterestMappings.json")
	var data []InteresetMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
