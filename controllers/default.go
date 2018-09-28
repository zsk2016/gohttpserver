package controllers

import (
	"gopkg.in/macaron.v1"
)

type Context struct {
	*macaron.Context
}

func Default(ctx *macaron.Context) {
	ctx.HTML(200, "pages/index")
}
