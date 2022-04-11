package main

import (
	"github.com/eduumach/biblioteca-backend/app"
	"os"
)

func main() {
	a := app.App{}
	a.Initialize(os.Getenv("DATABASE_URL"))

	a.Run(":" + os.Getenv("PORT"))
}
