package webcamserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	mux_context "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jalasoft/go-webcam"
	"github.com/jalasoft/go-webcam-server/params"
)

//go:generate go-bindata -pkg webcamserver -prefix assets assets/...

type camera_info struct {
	Name   string `json:"name"`
	Device string `json:device`
}

var parameters params.Params
var detectedDevicesInfo []camera_info

type RequestContextKey string

var cameraInfoContextKey RequestContextKey = "CAMERA_INFO_CONTEXT_KEY"

func StartServer() {

	parameters = getParameters()
	log.Printf("starting server on port %d", parameters.Port)

	detectedDevicesInfo = detectDevices()

	if len(detectedDevicesInfo) == 0 {
		log.Printf("No device detected. Exiting.")
		os.Exit(3)
	}

	log.Printf("Detected devices: %v\n", detectedDevicesInfo)

	rootRouter := mux.NewRouter()
	rootRouter.HandleFunc("/js/{resource:.+}", StaticContentHandler)
	rootRouter.HandleFunc("/", WebIndexHandler)
	restRouter := rootRouter.PathPrefix("/camera").Subrouter()
	restRouter.HandleFunc("/", allDevicesHandler).Methods("GET")

	constantsRouter := rootRouter.PathPrefix("/v4l2").Subrouter()
	constantsRouter.HandleFunc("/cap", AllCapabilitiesHandler).Methods("GET")
	constantsRouter.HandleFunc("/pixfmt", AllPixelFormatsHandler).Methods("GET")

	cameraRouter := restRouter.PathPrefix("/{camera}").Subrouter()
	cameraRouter.Use(cameraInfoInContextMiddleware)
	cameraRouter.HandleFunc("/", deviceInfoHandler).Methods("GET")
	cameraRouter.HandleFunc("/cap", deviceCapabilityHandler).Methods("GET")
	cameraRouter.HandleFunc("/stream", streamWebsocketHandler).Methods("GET")
	cameraRouter.HandleFunc("/stream2", streamHttpHandler).Methods("GET")

	//router.HandleFunc("/{name}", cameraHandler)
	//router.HandleFunc("/{name}/snapshot", snapshotHandler)
	//router.HandleFunc("/{name}/stream/web", streamWebIndexHandler)
	//router.HandleFunc("/{name}/stream/web/{res:[a-zA-Z0-9/\\.]+}", streamWebResourceHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", parameters.Port),
		Handler: rootRouter,
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

func getParameters() params.Params {
	par, err := params.ParseParams()

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	return par
}

func detectDevices() []camera_info {

	devices, err := webcam.SearchVideoDevices()

	if err != nil {
		log.Fatalf("Cannot get info about available devices: %v", err)
		os.Exit(2)
	}

	infos := make([]camera_info, 0, len(devices))

	for _, deviceFile := range devices {
		infos = append(infos, camera_info{
			Name:   filepath.Base(deviceFile),
			Device: deviceFile,
		})
	}

	return infos
}

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

func cameraInfoInContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		name, ok := extractVariable("camera", r)

		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No camera specified."))
			return
		}

		info, ok := cameraInfoByName(name)

		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("No camera '%s' found.", name)))
			return
		}

		mux_context.Set(r, cameraInfoContextKey, info)

		next.ServeHTTP(w, r)
	})
}

func extractVariable(name string, request *http.Request) (string, bool) {
	vars := mux.Vars(request)
	value, ok := vars[name]
	return value, ok
}

func cameraInfoByName(name string) (camera_info, bool) {
	for _, device := range detectedDevicesInfo {
		if device.Name == name {
			return device, true
		}
	}

	return camera_info{}, false
}
