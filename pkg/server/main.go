package server

// ref https://www.alexedwards.net/blog/serving-static-sites-with-go
// https://stackoverflow.com/questions/55996263/how-to-serve-static-files-for-all-pages-not-just-a-few
import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
  "encoding/json"
  "math/rand"
  "strconv"

	"goSite/pkg/posts"
	"goSite/pkg/jokes"
)

type joke1 struct {
        Joke string `json:"joke"`
}

var tpl *template.Template

//var projects []string ("one", "two")

func init() {
	tpl = template.Must(template.ParseGlob("tmpl/*.gohtml"))
}

func StartApp() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	router.HandleFunc("/", index)
	router.HandleFunc("/about", about)
	router.HandleFunc("/projects", projects)
	router.HandleFunc("/projects/{project}", project)
	router.HandleFunc("/resume", resume)
  router.HandleFunc("/api/joke/{id}", jokeByID)
  router.HandleFunc("/api/joke", joke)
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	http.ListenAndServe(":8080", router)
}

func NotFound(res http.ResponseWriter, req *http.Request) { // a * before http.Request
	err := tpl.ExecuteTemplate(res, "404.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func about(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
	if err != nil {
		log.Println("template didn't execute: ", err)
		NotFound(res, req)
	}
}

func projects(res http.ResponseWriter, req *http.Request) {
	projects := posts.Projects
	err := tpl.ExecuteTemplate(res, "projects.gohtml", projects)
	if err != nil {
		log.Println("template didn't execute: ", err)
		NotFound(res, req)
	}
}

func project(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	//id := params["project"]
	err := tpl.ExecuteTemplate(res, params["project"]+".gohtml", nil)
	if err != nil {
		log.Println("template didn't execute: ", err)
		NotFound(res, req)
		//tpl.ExecuteTemplate(res, "404.gohtml", nil)
	}
}

func resume(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "resume.gohtml", nil)
	if err != nil {
		log.Println("template didn't execute: ", err)
		NotFound(res, req)
	}
}

func jokeByID(res http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        id := vars["id"]

        j := &joke1 {
                Joke: jokes.Jokes[id],
        }
        jm, _ := json.Marshal(j)
        fmt.Fprintf(res, string(jm))
        //fmt.Fprintf(res, fmt.Sprintf("joke: %v", jokes.Jokes[id]))
}

func joke(res http.ResponseWriter, req *http.Request) {
        jokeID := strconv.Itoa(rand.Intn(len(jokes.Jokes)))

        j := &joke1 {
                Joke: jokes.Jokes[jokeID],
        }
        jm, _ := json.Marshal(j)
        fmt.Fprintf(res, string(jm))
}
