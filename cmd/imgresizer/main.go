package main

import (
	"fmt"

	"github.com/crazy-genius/imgresizer/internal/imgresizer/configuration"
	"github.com/crazy-genius/imgresizer/internal/imgresizer/http"
)

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {
	fmt.Println("Hello to resizer!")

	srv := http.NewServer(configuration.Configuration{
		Host:               "localhost",
		Port:               8080,
		EnableHTTPS:        false,
		EnableImageStorage: false,
		EnableRateLimit:    true,
		MaxConnections:     10,
		AllowedHosts:       []string{"*"},
	})

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
