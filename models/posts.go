package models

import (
	"errors"

	"goSite/pkg/posts"
	"goSite/views"
)

func PublishedPosts() views.Data {
	var vd views.Data
	vd.Yield = posts.Projects
	return vd
}

func PostExists(fileName string) error {
	published := posts.Projects
	if _, ok := published[fileName]; ok {
		return nil
	}
	return errors.New("Post not found")
}
