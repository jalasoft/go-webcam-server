package camserver

import (
	"log"
	"net/http"
)

func streamWebHandler(writer http.ResponseWriter, request *http.Request) {

	log.Printf("Prisel request....")
	/*
		ass, err := Asset("index.html")

		if err != nil {
			logAndWriteResponse("Canot find index.html", err, writer)
			return
		}

		log.Printf("index.html loaded")
		writer.Write(ass)*/
	writer.Write([]byte("Ahojky...."))
}
