package main

import "github.com/go-martini/martini"
import "thumbnailer/nailer"

func main() {
  m := martini.Classic()
  m.Get("/resize/:size", func(params martini.Params) string {

  	url := "http://localhost:3000/storms.jpg"

    return "<html><body>" +
						"<h1>Original</h1>" +
						"<img src='" + url + "'>" +
						"<h1>Thumbnailed</h1>" +
						"<img src='" + nailer.Process(url, params["size"]) + "'>" +
						"</body></html>"
  })
  m.Run()
}
