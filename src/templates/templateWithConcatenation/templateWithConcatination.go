package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Niteesh Kumar Dubey"
	str := fmt.Sprint(
		`
		Hello my name is` + name + `and Iam a male guy
		`,
	)

	nf, err := os.Create("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
