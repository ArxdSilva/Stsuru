package main

import (
	"github.com/iris-contrib/template/django"
	"github.com/kataras/iris"
)

func main() {

	iris.Render.Template(django.New()).Directory("./templates", ".html")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("mypage.html", map[string]interface{}{"username": "iris", "is_admin": true})
	})

	iris.Listen(":8080")
}
