package main

import (
	"fmt"
	"log"
	"net/http"

	"BackendGo/src/routers"
)

func main() {
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("aplikasi berjalan")
	
	if err := http.ListenAndServe(":8080", mainRoute); err != nil {
		log.Fatal("aplikasi gagal dijalankan")
	}

}