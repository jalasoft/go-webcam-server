package webcamserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/jalasoft/go-webcam"
)

type camera_capability struct {
	Camera       string   `json:"camera"`
	Capabilities []string `json:"capabilities"`
}

func deviceCapabilityHandler(writer http.ResponseWriter, request *http.Request) {

	cameraInfo := context.Get(request, cameraInfoContextKey).(camera_info)

	caps, err := readCapabilities(cameraInfo)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Cannot get info about capabilities for camera '%s': %v", cameraInfo.Name, err)))
		return
	}

	cap := camera_capability{
		Camera:       cameraInfo.Name,
		Capabilities: caps,
	}

	payload, err := json.MarshalIndent(cap, "", "  ")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Cannot create JSON response of camera '%s' capabilities: %v", cameraInfo.Name, err)))
		return
	}

	writer.Write(payload)
}

func readCapabilities(camera camera_info) ([]string, error) {

	device, err := webcam.OpenVideoDevice(camera.Device)

	defer func() {
		if err := device.Close(); err != nil {
			log.Printf("Cannot close device %s: %v\n", camera.Name, err)
		}
	}()

	if err != nil {
		return nil, err
	}

	capStrings := []string{}

	capabilities := device.Capabilities()

	for _, capability := range capabilities.AllCapabilities() {
		capStrings = append(capStrings, capability.Name)
	}

	return capStrings, nil
}
