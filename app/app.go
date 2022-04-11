package app

import (
	"github.com/eduumach/biblioteca-backend/app/controller"
	"github.com/eduumach/biblioteca-backend/app/model"
	"log"
	"net/http"
)

type App struct {
}

func (a *App) Initialize(database string) {

	db := model.DB{}
	db.Initialize(database)

	routers := controller.Routers{}
	routers.Initialize()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, controller.Routers{}.Router))
}
