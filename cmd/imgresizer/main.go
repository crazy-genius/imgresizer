package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/crazy-genius/imgresizer/internal/imgresizer/configuration"
	"github.com/crazy-genius/imgresizer/internal/imgresizer/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello to resizer!")

	if len(os.Args) < 2 {
		log.Fatal("No configuration file provided")
	}

	cfgArgument := os.Args[1]
	if !strings.Contains(cfgArgument, "-c=") {
		log.Fatal("No configuration file provided")
	}
	cfgArgument = strings.ReplaceAll(cfgArgument, "-c=", "")

	cfg, err := configuration.LoadConfiguration(cfgArgument)
	check(err)

	srv := http.NewServer(*cfg)

	srv.StartAndListenSignals()
}

// func resizeFromFile() {
// 	f, err := os.Open("./assets/IMG_7917.jpg")
// 	check(err)

// 	rs := resizer.NewResizer()
// 	newFile, _ := rs.Resize(bufio.NewReader(f), resizer.ResizeConfig{
// 		Dimenstions: resizer.Dimenstions{
// 			Height: 300,
// 			Width:  300,
// 		},
// 		Quality: 95,
// 	})
// 	check(f.Close())

// 	f, err = os.OpenFile("./assets/IMG_7917_new.jpg", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
// 	check(err)
// 	check(f.Close())
// 	_, err = f.Write(newFile)
// 	check(err)
// }
