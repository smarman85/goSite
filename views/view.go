package views

import (
  "io"
  "bytes"
  "errors"
  "net/url"
  "net/http"
  "html/template"
  "path/filepath"
  //"goSite/context"
)

var (
  LayoutDir string = "views/layouts/"
  TemplateDir string = "views/"
  TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
  addTemplatePath(files)
  addTemplateExt(files)
  files = append(files, layoutFiles()...)
  t, err := template.New("").Funcs(template.FuncMap{
    "csrfField": func() (template.HTML, error) {
      return "", errors.New("csrfField is not implemented")
    },
    "pathEscape": func(s string) string {
      return url.PathEscape(s)
    },
  }).ParseFiles(files...)
  if err != nil {
    panic(err)
  }
  return &View{
    Template: t,
    Layout: layout,
  }
}

type View struct {
  Template *template.Template
  Layout string
}

func addTemplatePath(files []string) {
  for i, f := range files {
    files[i] = TemplateDir + f
  }
}

func addTemplateExt(files []string) {
  for i, f := range files {
    files[i] = f + TemplateExt
  }
}

func layoutFiles() []string {
  files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
  if err != nil {
    panic(err)
  }
  return files
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  v.Render(w, r, nil)
}

func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
  w.Header().Set("Content-Type", "text/html")
  var vd Data
  switch d := data.(type) {
  case Data:
    vd = d
  default:
    vd = Data{
      Yield: data,
    }
  }
  if alert := getAlert(r); alert != nil {
    vd.Alert = alert
    clearAlert(w)
  }
  //vd.User = context.User(r.Context())
  var buf bytes.Buffer
  io.Copy(w, &buf)
}
