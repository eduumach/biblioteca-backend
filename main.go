package main

import (
	"os"
)

func main() {
	a := App{}
	a.Initialize(os.Getenv("DATABASE_URL"))

	a.Run(os.Getenv("PORT_APP"))
}
