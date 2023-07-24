package main

import "github.com/goellavish10/skillbee-assignment/app"

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
