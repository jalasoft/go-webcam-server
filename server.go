package webcamserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jalasoft/go-webcam-server/params"
)

//go:generate go-bindata -pkg webcamserver -prefix assets assets/...

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

	parameters = par

	log.Printf("starting server on port %d", parameters.Port)

	router := mux.NewRouter()

	//router.HandleFunc("/camera/stream/web/js/logic.js", streamWebResourceHandler)

	router = router.PathPrefix("/camera").Subrouter()
	router.HandleFunc("/", allCamerasHandler)
	router.HandleFunc("/{name}", cameraHandler)
	router.HandleFunc("/{name}/snapshot", snapshotHandler)
	router.HandleFunc("/{name}/stream/web", streamWebIndexHandler)
	router.HandleFunc("/stream/web/{res:[a-zA-Z0-9/\\.]+}", streamWebResourceHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", parameters.Port),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	<-signals

	log.Printf("Shutting down server")
	wait, _ := time.ParseDuration("5s")
	ctx, cancel := context.WithTimeout(context.Background(), wait)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error occurred during closing the server: %v", err)
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

func extractVariable(name string, request *http.Request) string {
	vars := mux.Vars(request)
	return vars[name]
}
