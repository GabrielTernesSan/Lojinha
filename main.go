package main

import (
	"net/http"

	routes "github.com/GabrielTernesSan/Lojinha/Routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
