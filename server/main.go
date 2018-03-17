package main

const host = "localhost:3000"

func main() {

	app := NewServer()
	app.listen(host)
}
