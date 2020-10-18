package main

import (
	"github.com/takutakukatokatojapan/image_analysis_api/injector"
	"log"
)

func main() {
	if err := injector.Run(); err != nil {
		log.Fatal(err)
	}
}
