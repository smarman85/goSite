package controllers

import "goSite/views"

func NewStatic() *Static{
  return &Static{
    Home: views.NewView(
      "elements", "static/home",
    ),
    About: views.NewView(
      "elements", "static/about",
    ),
    NotFound: views.NewView(
      "elements", "static/404",
    ),
  }
}

type Static struct {
  Home *views.View
  About *views.View
  NotFound *views.View
}
