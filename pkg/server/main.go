package server

// ref https://www.alexedwards.net/blog/serving-static-sites-with-go
// https://stackoverflow.com/questions/55996263/how-to-serve-static-files-for-all-pages-not-just-a-few
import (
	"github.com/gorilla/mux"
	"net/http"

	"goSite/pkg/handlers"
)

func StartApp() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/about", handlers.About)
	router.HandleFunc("/projects", handlers.Projects)
	router.HandleFunc("/projects/{project}", handlers.Project)
	router.HandleFunc("/resume", handlers.Resume)
  router.HandleFunc("/api/joke/{id}", handlers.JokeByID)
  router.HandleFunc("/api/joke", handlers.Joke)
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	http.ListenAndServe(":8080", router)
}
