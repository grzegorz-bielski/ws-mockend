package main

import (
	"fmt"
)

const port = ":3000"

func main() {

	app := NewServer()
	app.listen(port)

	fmt.Println("Listening on" + port)
}
