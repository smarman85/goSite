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
  pageContent := ""

  err := models.PostExists(projectName)
  if err != nil {
    //http.Redirect(w, r, "/posts", http.StatusTemporaryRedirect)
    pageContent = "static/404"
  } else {
    pageContent = fmt.Sprintf("posts/%s", projectName)
  }
  postView := views.NewView(
    "elements", pageContent,
  )
  postView.Render(w, r, postView)
}
