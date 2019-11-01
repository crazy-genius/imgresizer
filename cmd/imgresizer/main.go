package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/crazy-genius/imgresizer/internal/imgresizer/resizer"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello to resizer!")

	f, err := os.Open("./assets/IMG_7917.jpg")
	check(err)

	rs := resizer.NewResizer()
	newFile := rs.Resize(bufio.NewReader(f), resizer.ResizeConfig{
		Height: 300,
		Width:  300,
	})
	check(f.Close())

	f, err = os.OpenFile("./assets/IMG_7917_new.jpg", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	check(err)
	check(f.Close())
	_, err = f.Write(newFile)
	check(err)
}
