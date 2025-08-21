package main

import (
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/cmd"
)

func main() {	
	mux := http.NewServeMux()
	cmd.Serve(mux)
}


