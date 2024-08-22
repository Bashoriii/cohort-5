package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type Customers struct {
	ID       int
	Email    string
	Password string
	Address  string
	Job      string
	Reason   string
}

var customers = []Customers{
	{ID: 1, Email: "ujang@gmail.com", Password: "rahasia1", Address: "Bekasi", Job: "Athlete", Reason: "Your smile"},
	{ID: 2, Email: "dadang@gmail.com", Password: "rahasia2", Address: "Bekasi", Job: "Engineer", Reason: "Your smile"},
	{ID: 3, Email: "budi@gmail.com", Password: "rahasia3", Address: "Bekasi", Job: "Security", Reason: "Your smile"},
	{ID: 4, Email: "desy@gmail.com", Password: "rahasia4", Address: "Bekasi", Job: "Admin", Reason: "Your smile"},
	{ID: 5, Email: "malik@gmail.com", Password: "rahasia5", Address: "Bekasi", Job: "Engineer", Reason: "Your smile"},
	{ID: 6, Email: "sandi@gmail.com", Password: "rahasia6", Address: "Bekasi", Job: "Engineer", Reason: "Your smile"},
}

func HandleLoginFail(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		tpl, err := template.ParseFiles("sesi-5/error-page.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, nil)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func LoginCust(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")

		for _, customer := range customers {
			if customer.Email == email && customer.Password == password {
				http.Redirect(w, req, fmt.Sprintf("/profile?email=%s", email), http.StatusSeeOther)
				return
			}
		}

		http.Redirect(w, req, "/login-failed", http.StatusSeeOther)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func GetCustomers(w http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	// if req.Method == "GET" {
	// 	json.NewEncoder(w).Encode(customers)
	// 	return
	// }

	if req.Method == "GET" {
		tpl, err := template.ParseFiles("sesi-5/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, customers)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func ProfilePage(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		email := req.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email not provided", http.StatusBadRequest)
			return
		}

		var customer *Customers
		for _, c := range customers {
			if c.Email == email {
				customer = &c
				break
			}
		}

		if customer == nil {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}

		tpl, err := template.ParseFiles("sesi-5/profile.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, customer)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func CreateCust(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")

		newCustomers := Customers{
			ID:       len(customers) + 1,
			Email:    email,
			Password: password,
		}

		customers = append(customers, newCustomers)
		json.NewEncoder(w).Encode(customers)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
