package webcamserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/jalasoft/go-webcam"
)

func frameSizeEndpoint(writer http.ResponseWriter, request *http.Request) {

	cameraInfo := context.Get(request, cameraInfoContextKey).(camera_info)

	frmtype, ok := extractVariable("frametype", request)

	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "Type of frame not specified")
		return
	}

	pixfmtparam, ok := request.URL.Query()["pixfmt"]

	if !ok {
		pixfmtparam = []string{"V4L2_PIX_FMT_MJPEG"}
	}

	pixfmt, ok := webcam.PixelFormatFromString(pixfmtparam[0])

	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Unknown pixel format: %v", pixfmtparam[0])
		return
	}

	switch frmtype {
	case "stepwise":
		readStepwiseFrameSizes(writer, cameraInfo, pixfmt)
	case "discrete":
		readDiscreteFrameSizes(writer, cameraInfo, pixfmt)
	default:
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Unexpected type of frame: %s", frmtype)
	}
}

func readDiscreteFrameSizes(writer http.ResponseWriter, camera camera_info, pixfmt webcam.PixelFormat) {
	device, err := webcam.OpenVideoDevice(camera.Device)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Cannot open device: %v", err)
		return
	}

	defer func() {
		if err := device.Close(); err != nil {
			fmt.Fprintf(writer, "Cannot close device %s: %v\n", camera.Name, err)
		}
	}()

	framesizes := device.FrameSizes()
	frmsizes, err := framesizes.Discrete(pixfmt)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Cannot read discrete frame sizes of type %s: %v", pixfmt.Name, err)
		return
	}

	b, err := json.MarshalIndent(frmsizes, "", "  ")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Cannot create payload: %v", err)
		return
	}

	writer.Write(b)
}

func readStepwiseFrameSizes(writer http.ResponseWriter, camera camera_info, pixfmt webcam.PixelFormat) {
	device, err := webcam.OpenVideoDevice(camera.Device)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Cannot open device: %v", err)
		return
	}

	defer func() {
		if err := device.Close(); err != nil {
			fmt.Fprintf(writer, "Cannot close device %s: %v\n", camera.Name, err)
		}
	}()

	framesizes := device.FrameSizes()
	frmsizes, err := framesizes.Stepwise(pixfmt)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Cannot read stepwise frame sizes of type %s: %v", pixfmt.Name, err)
		return
	}

	b, err := json.MarshalIndent(frmsizes, "", "  ")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Cannot create payload: %v", err)
		return
	}

	writer.Write(b)
}
