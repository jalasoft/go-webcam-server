package camserver

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jalasoft/go-webcam-server/params"
)

//go:generate go-bindata -pkg camserver -prefix "assets/" ./assets/...

type cameraInfo struct {
	Driver  string `json:"driver"`
	Card    string `json:"card"`
	Businfo string `json:"bus_info"`
}

var parameters params.Params

func StartServer() {

	par, err := params.ParseParams()

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	io
	parameters = par

	log.Printf("starting server on port %d", parameters.Port)
	log.Printf("Startujuuuuuuuu")

	router := mux.NewRouter()

	router.Methods("GET")
	router.HandleFunc("/camera/", allCamerasHandler)
	router.HandleFunc("/camera/{name}", cameraHandler)
	//router.HandleFunc("/camera/{name}/snapshot", snapshotHandler)

	router.HandleFunc("/camera/{name}/stream/web", streamWebHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", parameters.Port),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//-------------------------------------------------------------------------------------
//UTILITY FUNCTIONS
//-------------------------------------------------------------------------------------

func logAndWriteResponse(m string, err error, writer http.ResponseWriter) {
	var message string
	if err != nil {
		message = fmt.Sprintf("%v: %v\n", m, err)
	} else {
		message = m
	}

	log.Printf(message)
	writer.Write([]byte(message))
}
