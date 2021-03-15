// ref https://www.alexedwards.net/blog/serving-static-sites-with-go
// https://stackoverflow.com/questions/55996263/how-to-serve-static-files-for-all-pages-not-just-a-few
package main

import (
	"goSite/pkg/server"
)

func main() {
	server.StartApp()
}
