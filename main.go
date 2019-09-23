package main

import (
	"net/http"

	"github.com/vima91/GoPlayground/controllers"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
