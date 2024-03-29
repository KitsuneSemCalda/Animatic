package main

import (
	cli "Animatic/Cli"
	"log"
)

func main() {
	animeName, err := cli.GetUserInput("Enter the name of the anime you want to download")

	if err != nil {
		log.Println("Occurred an unknown error: ", err.Error())
	}

	log.Println(*animeName)
}
