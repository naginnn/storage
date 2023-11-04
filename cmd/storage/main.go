package main

import (
	"fmt"
	"github.com/naginnn/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("Hello"))

	if err != nil {
		log.Fatal(err)
	}

	downloadFile, err := st.GetByID(file.ID)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it works!", downloadFile)
}
