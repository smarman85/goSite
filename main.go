// ref https://www.alexedwards.net/blog/serving-static-sites-with-go
// https://stackoverflow.com/questions/55996263/how-to-serve-static-files-for-all-pages-not-just-a-few
package main

import (
  //"fmt"
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
)

func staticHTML(htmlTemplate string, resp http.ResponseWriter) {
  templ := template.Must(template.ParseFiles("tmpl/" + htmlTemplate + ".html"))
  templ.Execute(resp, nil)

}

func main() {
  router:= mux.NewRouter()

  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    staticHTML("index", w)
  })

  router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "About me")
    staticHTML("about", w)
  })

  router.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request){
    staticHTML("projects", w)
  })

  router.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request){
    staticHTML("resume", w)
  })

  //http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  http.ListenAndServe(":8088", router)
}
