package main

import (
	"compro/cmd"
	"log"
	"net/http"
)

func main()  {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Execute the Cobra command (which will start your main application server)
	cmd.Execute()
}
