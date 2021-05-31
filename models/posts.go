package models

import (
  "os"
  "errors"

  "goSite/views"
  "goSite/pkg/posts"
)

func PublishedPosts() views.Data {
  var vd views.Data
  vd.Yield = posts.Projects
  return vd
}

func PostExists(fileName string) error {
  _, err := os.Stat(fileName)
  if os.IsNotExist(err) {
    return errors.New("the post doesn't exist")
  }
  return nil
}
