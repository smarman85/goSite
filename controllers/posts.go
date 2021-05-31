package controllers

import(
  "net/http"

  "goSite/views"
  "goSite/pkg/posts"

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
  //publishedPosts := posts.Projects
  var vd views.Data
  vd.Yield = posts.Projects
  p.ShowView.Render(w, r, vd)
}

