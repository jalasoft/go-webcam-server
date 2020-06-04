package webcamserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func allDevicesHandler(writer http.ResponseWriter, request *http.Request) {

	b, err := json.MarshalIndent(detectedDevicesInfo, "", "  ")

	if err != nil {
		writer.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}

	writer.Write(b)
}