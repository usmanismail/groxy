package main

import (
	"log"

	"github.com/usmanismail/groxy/groxyapp"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Terminating application: ", r)

		}
	}()

	groxyapp_, err := groxyapp.New()
	if err == nil {
		groxyapp_.Start()
	} else {
		log.Fatal(err.Error())
	}
}
