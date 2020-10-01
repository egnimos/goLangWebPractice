package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprint(
		`
		Hello my name is ` + name + `and currently iam using the os library with args value
		`,
	)
	nf, err := os.Create("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
