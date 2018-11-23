package main

import(
	"fmt"
	"log"
	"dellemc/system-state-service/src/routes"
	"net/http"
)

func main(){

	r := routes.NewRouter()
	fmt.Println("Starting Server ..........")
	fmt.Println("Listening on port 8000 ..........")
	log.Fatal(http.ListenAndServe(":8000", r))
}	
