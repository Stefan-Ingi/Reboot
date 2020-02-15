package main

import (
	"encoding/json"
	"net/http"

	"github.com/BjornGudmundsson/Reboot/insurances"
	"github.com/BjornGudmundsson/Reboot/users"
)

func GetCookie(phone string) *http.Cookie {
	return &http.Cookie{
		Name:  "Reboot",
		Value: phone,
	}
}

func LoginForm(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers: Content-Type", "Authorization")
	j := json.NewDecoder(req.Body)
	var js users.User
	e := j.Decode(&js)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ph := js.Number
	e = users.CheckIfUserExists(ph)
	if e != nil {
		users.WriteUserToDB(ph)
	}
	e = users.LoginUser(ph)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		http.SetCookie(w, GetCookie(ph))
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	http.HandleFunc("/loginForm", LoginForm)
	http.HandleFunc("/addInsurance", insurances.AcceptInsurance)
	http.HandleFunc("/myInsurances", insurances.GetMyInsurances)
	http.HandleFunc("/search", insurances.SearchForInsurance)
	http.ListenAndServe(":8084", nil)
}