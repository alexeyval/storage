package main

import (
	"fmt"
	"github.com/alexeyval/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("it uploaded", file)

	fmt.Println("it is restored", restoredFile)
}
