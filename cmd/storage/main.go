package main

import (
	"fmt"

	"github.com/NevostruevK/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("it works", st)
}