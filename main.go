package main

import (
	"log"
	"os"

	"github.com/PaulaSalinoRibeiro/cli-with-go-lang/app"
)

func main() {

	appli := app.Gerar()

	if error := appli.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}
