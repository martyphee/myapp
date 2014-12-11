package controllers

import (
	"fmt"
	"github.com/revel/revel"
	models "myapp/app/models"
	//	"github.com/stripe/stripe-go"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// E.g finding an existing User
	var user *models.User

	user = models.DAO.GetFirst()

	fmt.Println(user)

	return c.Render(user)
}
