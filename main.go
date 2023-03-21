package main

import (
	"github.com/RickHPotter/flutter_rest_api/initialisers"
	"github.com/RickHPotter/flutter_rest_api/routers"
)

func init() {
	initialisers.LoadEnv()
	initialisers.PrepareDatabase()
}

// ! run with `compiledaemon --command="./flutter_rest_api"`
func main() {
	routers.LoadRouters()
}
