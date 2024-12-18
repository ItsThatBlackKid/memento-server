package main

import "os"

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("DB"))
	app.Run(":8080")
}
