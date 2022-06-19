package main

import (
	"BackendGo/src/configs/command"
	"log"
	"os"
)

func main() {
	// mainRoute, err := routers.New()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("aplikasi berjalan pada port 8080")

	// if err := http.ListenAndServe(":8080", mainRoute); err != nil {
	// 	log.Fatal("aplikasi gagal dijalankan")
	// }
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

}
