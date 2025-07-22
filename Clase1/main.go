package main

import(
	"fmt"
	"log"
	"net/http"
	"clase1/controller"
)

func main(){
	http.HandleFunc("/analizar", controller.AnalizarArchivoHandler)
	fmt.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}