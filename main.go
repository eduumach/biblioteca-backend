package main

import (
	"os"
)

func main() {
	a := App{}
	a.Initialize(os.Getenv("DATABASE_URL"))

	a.Run(":8000")
}
