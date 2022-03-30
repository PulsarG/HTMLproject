package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Person struct {
	Name                  string
	Age                   uint16
	Money                 int
	avg_grades, happiness float64
	Hobbis []string
}

func (u Person) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money: %d", u.Name, u.Age, u.Money)
}

/* func (u *Person) setNewName (newName string) {
	u.Name = newName
} */

func index(w http.ResponseWriter, r *http.Request) {
	Bob := Person{"Bob", 25, -50, 0.6, 0.8, []string{"First", "Second", "Third"} }
	//Bob.setNewName("Alex")
	//fmt.Fprintf(w, "index.html")

	tmpl, _ := template.ParseFiles("test.html")
	tmpl.Execute(w, Bob)
}

/* func secondPage(w http.ResponseWriter, r *http.Request) {
	Bob := Person{"Bob", 25, -50, 0.6, 0.8}
	tmpl2, _ := template.ParseFiles("secondpage.html")
	tmpl2.Execute(w, Bob)
} */

func handlRequest() {
	http.HandleFunc("/", index)
	//http.HandleFunc("/secondpage/", secondPage)
	http.ListenAndServe(":5500", nil)
}

// *******************************************************************************************

func main() {

	handlRequest()
}
