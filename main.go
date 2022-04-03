package main

import (
	"fmt"
	"html/template"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id                     uint16
	Title, Anons, FullText string
}

var posts = []Article{}

func Index(w http.ResponseWriter, r *http.Request) { // функция открытия главной страницы. Ниже повторный код для каждой страницы
	t, err := template.ParseFiles("html/index.html", "html/header.html", "html/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil {
		panic(err)
	}

	posts = []Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Anons, &post.Title, &post.FullText)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	t.ExecuteTemplate(w, "index", posts)
}

func SecondPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/secondpage.html", "html/header.html", "html/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "secondpage", nil)
}

func SaveArticle(w http.ResponseWriter, r *http.Request) { // функция добавления данных в БД
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprint(w, "Не все данные заполнены")
	} else {

		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')", title, anons, full_text))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// ***************************************************************************************************************************************************************************

func handlFunc() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	http.Handle("/fav/", http.StripPrefix("/fav/", http.FileServer(http.Dir("./fav/"))))
	http.HandleFunc("/", Index)
	http.HandleFunc("/secondpage", SecondPage)
	http.HandleFunc("/save_article", SaveArticle) // Первый аргумент - адресс в адресной строке, он же пишется как ссылка на кнопках и т.п. Второй аргумент - функция выполнения
	http.ListenAndServe(":5500", nil)
}

func main() {
	handlFunc()
}
