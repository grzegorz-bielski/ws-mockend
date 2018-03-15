package main

import (
	"fmt"
)

func main() {

	for _, data := range GetWSConfig() {
		fmt.Println(data.stringify())
	}
}
