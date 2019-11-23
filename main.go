package main

import (
	"net/http"

	"github.com/personal/course-go-starting/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":1212", nil)
}
