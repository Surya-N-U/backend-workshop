package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	ID      int64
	Name    string
	Email   string
	Phone   string
	Message string
}

func HandleTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	tmpl.Execute(w, data)
	return
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Handling GET requests
		fmt.Println("GET request received for /about")
		htmlContent, err := os.ReadFile("static/about.html")
		if err != nil {
			http.Error(w, "Failed to read HTML file", http.StatusInternalServerError)
			return
		}
		// Set the Content-Type header to indicate HTML content
		w.Header().Set("Content-Type", "text/html")
		// Send the HTML content as the response
		w.Write(htmlContent)
	case http.MethodPost:
		// Handling POST requests
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}
		// Accessing form data
		name := r.Form.Get("name")
		email := r.Form.Get("email")
		phone := r.Form.Get("phone")
		message := r.Form.Get("message")
		// Responding with a personalized message
		fmt.Println("Hello! POST request received for /about")
		fmt.Println(name + "\n" + email + "\n" + phone + "\n" + message + "\n")
		// Set the Content-Type header to indicate HTML content
		// Send the HTML content as the response
		tm, err := template.ParseFiles("received.html")
		if err != nil {
			http.Error(w, "Failed to parse HTML file", http.StatusInternalServerError)
			return
		}
		data1 := Data{
			Name:    name,
			Email:   email,
			Phone:   phone,
			Message: message,
		}
		DbConnect(data1)
		data2, err := GetData()
		if err != nil {
			http.Error(w, "Failed to get data", http.StatusInternalServerError)
			return
		}
		HandleTemplate(w, tm, data2)
	default:
		// Handling unsupported methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Registering the aboutHandler for the /about endpoint
	http.HandleFunc("/about", aboutHandler)

	// Starting the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
