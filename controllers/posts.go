package controllers

import(
  "fmt"
  "net/http"

  "goSite/views"
  "goSite/models"

  "github.com/gorilla/mux"
)


func NewPosts(r *mux.Router) *Posts{
  return &Posts{
    ShowView: views.NewView(
      "elements", "posts/index",
    ),
  }
}

type Posts struct {
  ShowView *views.View
}

func (p *Posts) ShowAll(w http.ResponseWriter, r *http.Request) {
  vd := models.PublishedPosts()
  p.ShowView.Render(w, r, vd)
}


func (p *Posts) Project(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  projectName := vars["post"]
  postView := views.NewView(
    "elements", fmt.Sprintf("posts/%s", projectName),
  )
  postView.Render(w, r, postView)
}
