package webcamserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jalasoft/go-webcam"
)

func AllCapabilitiesHandler(writer http.ResponseWriter, request *http.Request) {

	payload := make([]string, 0, len(webcam.AllCapabilities))

	for _, cap := range webcam.AllCapabilities {
		payload = append(payload, cap.Name)
	}

	writePayload(writer, payload)
}

func AllPixelFormatsHandler(writer http.ResponseWriter, request *http.Request) {

	payload := make([]string, 0, len(webcam.AllPixelFormats))

	for _, f := range webcam.AllPixelFormats {
		payload = append(payload, f.Name)
	}

	writePayload(writer, payload)
}

func writePayload(writer http.ResponseWriter, payload []string) {
	b, err := json.MarshalIndent(payload, "", "  ")

	if err != nil {
		writer.Write([]byte(fmt.Sprint(err)))
		return
	}

	writer.Write(b)
}
