package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"strconv"
	"time"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/go-rod/rod"
	//"github.com/go-rod/rod/lib/input"
	//"github.com/go-rod/rod/lib/proto"
	"github.com/gorilla/mux"
	//"github.com/go-rod/rod/lib/launcher"
)

var pwd, _ = os.Getwd()

type runStruct struct{
	Scope     []string
	UsersMax  int

}


func run(data runStruct){

	for i := 0; i < data.UsersMax; i++ {
		fmt.Printf("%T\n", data.Scope)
		fmt.Printf("%T\n", data.UsersMax)
		time.Sleep(60 * time.Second)
	}
	


}



func Home(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(pwd + "/template/index.html", pwd + "/template/menu/leftMenu.html", pwd + "/template/content/content.html", pwd + "/template/menu/mainMenu.html")
	fmt.Println(err)
	tmpl.ExecuteTemplate(w, "main", nil)

}




func HomePOST(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(pwd + "/template/index.html")
	tmpl.ExecuteTemplate(w, "main", nil)
		
}

func AddGet(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(pwd + "/template/index.html", pwd + "/template/menu/leftMenu.html", pwd + "/template/content/add.html", pwd + "/template/menu/mainMenu.html")
	fmt.Println(err)
	tmpl.ExecuteTemplate(w, "main", nil)

}

func AddPost(w http.ResponseWriter, r *http.Request) {
	users, _ := strconv.Atoi(r.FormValue("users"))
	scope := r.FormValue("scope")
	scopeArr := strings.Split(scope, ",")

	go run(runStruct{scopeArr, users})
	tmpl, err := template.ParseFiles(pwd + "/template/index.html", pwd + "/template/menu/leftMenu.html", pwd + "/template/content/add.html", pwd + "/template/menu/mainMenu.html")
	fmt.Println(err)
	tmpl.ExecuteTemplate(w, "main", nil)

}


func main() {

	//u := launcher.New().Set("user-data-dir", "path").Set("headless").Delete("--headless").MustLaunch()
	//browser := rod.New().MustConnect()

	// Even you forget to close, rod will close it after main process ends.
	//defer browser.MustClose()

	// Create a new page
	// page := browser.MustPage("http://127.0.0.1:9000/login")
	// page.MustElement("[name=Email]").MustInput("test@test.com")
	// page.MustElement("[name=Password]").MustInput("123456789Qq").MustType(input.Enter)
	

	// sess = page.MustCookies()
	// fmt.Println(sess)
	fmt.Println("Server start")
	r := mux.NewRouter()

	//Домашняя страница
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/", HomePOST).Methods("POST")
	r.HandleFunc("/add", AddGet).Methods("GET")
	r.HandleFunc("/add", AddPost).Methods("POST")

	//Подключаем структуры
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)
	// start the server on port 8000
	log.Fatal(http.ListenAndServe("127.0.0.1:2051", r))

}
