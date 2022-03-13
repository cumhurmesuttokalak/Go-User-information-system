package handlers

import (
	"encoding/json"
	"net/http"

	. "../dataloaders"
	. "../models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
func Handler(wr http.ResponseWriter, r *http.Request) {
	//Page Nesnesi
	sayfa := Sayfa{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	//Data Loaders
	users := LoadUsers()
	interests := LoadInterests()
	interestsMappings := LoadInterestMappings()
	//işlem
	var newUsers []User
	for _, user := range users {
		for _, interestsMapping := range interestsMappings {
			if user.ID == interestsMapping.UserID {
				for _, interest := range interests {
					if interestsMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
		//fmt.Println(user.FirstName)
	}
	viewModel := UserViewModel{Sayfa: sayfa, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	wr.Write([]byte(data))
}
