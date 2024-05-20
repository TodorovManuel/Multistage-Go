package main
import (
	"fmt"
	"net/http"
	"text/template"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("templates/index.html")
	template.Execute(rw, nil)
}

func main() {
	http.HandleFunc("/", Index)

	//Creación del servidor
	fmt.Println("El servidor esta corriedno en el puerto 3000")
	http.ListenAndServe("0.0.0.0:3000", nil)
}