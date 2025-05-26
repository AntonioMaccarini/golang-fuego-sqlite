package main

import (
	server "golang-fuego-sqlite/lib"
)

func main() {

	err := server.NewGrafanaStoreServer().Run()
	if err != nil {
		panic(err)
	}
}
