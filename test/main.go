package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	uuidstr := uuid.NewV4().String()

	fmt.Println(uuidstr)
}


