package models

import (
  "goSite/views"
  "goSite/pkg/posts"
)

func PublishedPosts() views.Data {
  var vd views.Data
  vd.Yield = posts.Projects
  return vd
}
