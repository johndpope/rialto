package main

import (
	"github.com/go-macaron/pongo2"
	macaron "gopkg.in/macaron.v1"
)

func main() {
	m := macaron.New()
	m.Use(pongo2.Pongoer(pongo2.Options{
		// Directory to load templates. Default is "templates".
		Directory: "/webapp",
		// Extensions to parse template files from. Defaults are [".tmpl", ".html"].
		Extensions: []string{".tmpl", ".html"},
		// Appends the given charset to the Content-Type header. Default is "UTF-8".
		Charset: "UTF-8",
		// Allows changing of output to XHTML instead of HTML. Default is "text/html".
		HTMLContentType: "text/html",
	}))
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "HelloServer"
		ctx.HTML(200, "index")
	})
	m.Run()
}
