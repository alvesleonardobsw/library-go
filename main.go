package main

import (
	"github.com/alvesleonardobsw/library-go/controller"
)

func main() {
	r := controller.NewRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
