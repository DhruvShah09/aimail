package testing

import (
	"fmt"
	"log"
	"os"
)

func writeResponseToFile(z string) {
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(z)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
