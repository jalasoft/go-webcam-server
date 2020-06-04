package webcamserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jalasoft/go-webcam"
)

type camera_detail struct {
	Name    string `json:"name"`
	File    string `json:"file"`
	Driver  string `json:"driver"`
	Card    string `json:"card"`
	Businfo string `json:"bus_info"`
	Version uint32 `json:"version"`
}

func deviceInfoHandler(writer http.ResponseWriter, request *http.Request) {


	deviceName, ok := extractVariable("name", request)

	if !ok {
		writer.Write([]byte("No device specified."))
		return
	}

	deviceInfo, ok := deviceInfoByName(deviceName)

	if !ok {
		writer.Write([]byte(fmt.Sprintf("No device named '%s'", deviceName)))
		return
	}

	cameraDetail, err := readCameraDetail(deviceInfo)
	
	if err != nil {
		writer.Write([]byte(fmt.Sprintf("An error occurred: %v", err)))
		return
	}

	b, err := json.MarshalIndent(cameraDetail, "", "  ")

	if err != nil {
		writer.Write([]byte(fmt.Sprintf("%v", err)))
	}

	writer.Write(b)
}

func readCameraDetail(camInfo camera_info) (camera_detail, error) {

	device, err := webcam.OpenVideoDevice(camInfo.Device)

	info := camera_detail{}
	info.Name = camInfo.Name
	info.File = camInfo.Device

	defer func() {
		if err := device.Close(); err != nil {
			log.Printf("Cannot close device %s: %v\n", camInfo.Name, err)
		}
	}()

	if err != nil {
		info.Driver = fmt.Sprintf("cannot load: %v", err)
		return info, err
	}

	cap := device.Capability()

	info.Driver = trim(cap.Driver())
	info.Card = trim(cap.Card())
	info.Businfo = trim(cap.BusInfo())
	info.Version = cap.Version()

	fmt.Printf("'%v'\n", info.Driver)

	return info, nil
}

func trim(value string) string {
	return strings.Trim(value, string('\u0000'))
}