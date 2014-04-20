package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"thumbnailer/nailer"
)

type ImageUrl struct {
  Url 		string `form:"url"`
  Width		string `form:"width"`
}

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  m.Get("/resize", func(r render.Render) {
  	r.HTML(200, "size", nil)
  })

  m.Post("/resize", binding.Form(ImageUrl{}), func(imageUrl ImageUrl, r render.Render) {
  	if imageUrl.Url == "" {
  		// TODO: Return an error msg
  	} else if imageUrl.Width == "" {
  		// TODO: Return an error msg
  	} else {
    	nailer.Thumbnail(imageUrl.Url, imageUrl.Width)
    }
    r.HTML(200, "size", nil)
  })

  m.Run()
}
