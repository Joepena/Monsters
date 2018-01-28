package main

import (
	"log"

	//"github.com/joepena/monsters/actions"
	"Monsters/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
