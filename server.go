package main

import "github.com/go-martini/martini"
import "thumbnailer/nailer"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "<html><body>" +
						"<h1>Original File</h1>" +
						"<img src='storms.jpg'>" +
						"<h1>Thumbnailed</h1>" +
						"<img src='" + nailer.Process("public/storms.jpg") + "'>" +
						"</body></html>"
  })
  m.Run()
}
