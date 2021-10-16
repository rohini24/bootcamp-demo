package controller

import (
	"bootcamp_demo/core"
	"bootcamp_demo/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	core.Controller
}

func (a *UserController) Show(ctx *gin.Context) {
	result, user := new(models.User).GetById(464) // @TODO: Get this from request
	a.Json(ctx, result, user)
}

// @TODO: Create a store api for user
